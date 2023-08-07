

func main() {
	appCfg, err := config.GetConfig(*envPtr)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get app config")
	}
	log.Info().Str("Address: currently running", appCfg.GetRedisConfig().Addr).Str("Password", appCfg.GetRedisConfig().Password).Str("Port", appCfg.GetPort()).Str("Username", appCfg.GetRedisConfig().Username).Msg("Configuration check")
	rdb := redis.NewClient(&redis.Options{
		Addr:     appCfg.GetRedisConfig().Addr,
		Password: appCfg.GetRedisConfig().Password,
		Username: appCfg.GetRedisConfig().Username,
	})
}