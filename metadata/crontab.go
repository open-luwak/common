package metadata

type CrontabConfig struct {
	CronJobs []*CronJob `toml:"crontab,omitempty"`
}

type CronJob struct {
	JobName   string `toml:"job_name"`
	JobCron   string `toml:"job_cron"`
	API       string `toml:"api"`
	APIParams any    `toml:"api_params,omitempty"`
}
