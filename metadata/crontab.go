package metadata

type CrontabInfo struct {
	JobName   string `toml:"job_name"`
	JobCron   string `toml:"job_cron"`
	API       string `toml:"api"`
	APIParams string `toml:"api_params"`
}
