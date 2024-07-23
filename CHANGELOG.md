* Added support of `YDB_OAUTH2_KEY_FILE` environment variable

## v0.4.0
* Added support `use_env_credentials` DSN parameter for use credentials from environ

## v0.3.1
* Added support of `YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS`, `YDB_STATIC_CREDENTIALS_{USER,PASSWORD,ENDPOINT}` environment variables
* Excluded auto-adding `WithInternalCA` on from option `environ.WithEnvironCredentials()` (because TLS certificates is not about credentials, it is about transport setting)
* Removed useless context argument from `environ.WithEnvironCredentials()`
* Changed default credentials (if not defined known environment variables) to anonymous credentials
