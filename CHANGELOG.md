## v0.3.1
* Added support of `YDB_SERVICE_ACCOUNT_KEY_CREDENTIALS`, `YDB_STATIC_CREDENTIALS_{USER,PASSWORD,ENDPOINT}` environment variables
* Excluded auto-adding `WithInternalCA` on from option `environ.WithEnvironCredentials()` (because TLS certificates is not about credentials, it is about transport setting)
* Removed useless context argument from `environ.WithEnvironCredentials()`
* Changed default credentials (if not defined known environment variables) to anonymous credentials
