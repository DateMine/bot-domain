package service_provider

import (
	"context"
	"github.com/DateMine/bot-domain/internal/config"
	"github.com/DateMine/bot-domain/pkg/models/offer_storage"
	"github.com/DateMine/bot-domain/pkg/repository"
	"github.com/DateMine/bot-domain/pkg/repository/pg/offer"
	"github.com/DateMine/bot-domain/pkg/repository/pg/parser"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type ServiceProvider struct {
	offerStorage     *offer_storage.OfferStorage
	offerRepository  repository.OfferRepository
	parserRepository repository.ParserRepository
	grpcConfig       config.GRPCConfig
	pgConfig         config.PGConfig
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) OfferStorage(ctx context.Context) *offer_storage.OfferStorage {
	if sp.offerStorage == nil {
		sp.offerStorage = offer_storage.NewOfferStorage(sp.OfferRepository(sp.PGConfig(), ctx))
	}

	return sp.offerStorage
}

func (sp *ServiceProvider) OfferRepository(pgConfig config.PGConfig, ctx context.Context) repository.OfferRepository {
	if sp.offerRepository == nil {
		config, err := pgxpool.ParseConfig(pgConfig.DSN())
		if err != nil {
			log.Printf("Unable to parse connection string: %v\n", err)
		}

		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			log.Printf("Unable to connect to database: %v\n", err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Printf("Unable to ping database: %v\n", err)
		}
		sp.offerRepository = offer.NewRepository(pool)
	}

	return sp.offerRepository
}

func (sp *ServiceProvider) ParserRepository(pgConfig config.PGConfig, ctx context.Context) repository.ParserRepository {
	if sp.parserRepository == nil {
		config, err := pgxpool.ParseConfig(pgConfig.DSN())
		if err != nil {
			log.Printf("Unable to parse connection string: %v\n", err)
		}
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			log.Printf("Unable to connect to database: %v\n", err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Printf("Unable to ping database: %v\n", err)
		}
		sp.parserRepository = parser.NewRepository(pool)
	}

	return sp.parserRepository
}

func (sp *ServiceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Printf("Unable to connect to database: %v\n", err)
			return nil
		}
		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *ServiceProvider) GrpcConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Printf("Unable to connect to database: %v\n", err)
			return nil
		}
		sp.grpcConfig = cfg
	}

	return sp.grpcConfig
}
