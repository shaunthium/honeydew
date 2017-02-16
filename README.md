# honeydew - Integrate Version Control with your environment
A lightweight CLI tool that automates Docker containerization, environment provisioning and deployment of your application.

Features:
- Clone entire environments and databases within seconds, using NetApp's cloud data replication and cloning features
- Provision custom environments, complete with database and environment configurations
- Automatically containerize your app using Docker, push it to an instance provisioned with the custom environment of your choosing and run it with a provisioned Docker volume
- Automatically teardown your environments when you're done, leaving it the way it was before
- An `api` mode which provides all the above functionality through a RESTful API
- Do all the above with single-touch, "one-click" commands

NOTE: Secrets are not stored in the repo. To run `honeydew`, create a `secrets.json` file with the following keys:
`netapp_base_url`, `netapp_key`, `instance_username` and `pem_file_path`, and set the path to the file as an environment variable with the key of `SECRETS_PATH`.
