package app

import (
	"github.com/jmoiron/sqlx"
	v1 "github/cntrkilril/dynamic-user-segmentation-golang/internal/controller/http/v1"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/infrastructure"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/service"
	"github/cntrkilril/dynamic-user-segmentation-golang/pkg/govalidator"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jackc/pgx/stdlib"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Run() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	atom := zap.NewAtomicLevel()
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			return
		}
	}(logger)

	l := logger.Sugar()
	atom.SetLevel(zapcore.Level(*cfg.Logger.Level))
	l.Infof("logger initialized successfully")

	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	f.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	f.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	l.Infof("fiber initialized successfully")

	// db
	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		l.Error(err)
		return
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			l.Error(err)
			return
		}
	}(db)

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime * time.Second)

	err = db.Ping()
	if err != nil {
		l.Error(err)
		return
	}

	l.Debug("Connected to PostgreSQL")

	// validator
	val := govalidator.New()

	// infrastructures
	segmentRepo := infrastructure.NewSegmentRepository(*db)
	usersSegmentsRepo := infrastructure.NewUsersSegmentsRepository(*db)
	registryGateway := infrastructure.NewPGRegistry(db)

	// services
	segmentService := service.NewSegmentService(registryGateway)
	usersSegmentsService := service.NewUsersSegmentsService(usersSegmentsRepo, segmentRepo)

	// controllers
	segmentHandler := v1.NewSegmentHandler(segmentService, val)
	usersSegmentsHandler := v1.NewUsersSegmentsHandler(usersSegmentsService, val)

	// groups
	apiGroup := f.Group("api")
	segmentGroup := apiGroup.Group("segment")

	segmentHandler.Register(segmentGroup)
	usersSegmentsHandler.Register(segmentGroup)

	go func() {
		err = f.Listen(net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port))
		if err != nil {
			l.Fatal(err.Error())
		}
	}()

	l.Debug("Started HTTP server")

	l.Debug("Application has started")

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	l.Info("Application has been shut down")

}
