# Running a local Docker build

In order to run the docker build locally, you must have the `.netrc` file configured locally.
See [netrc.md](netrc.md) for more.

## Running a build
To run a build, use the make command
```shell
make build
```

The first thing that will happen is you'll be prompted to log into `gcloud`.

After logging in, the build command will use the password from your `.netrc` and set 
the build arguments required to pull private Zenbusiness packages.

That this point you might see something with a success message:

```text
Step 17/18 : EXPOSE 50051 3001
 ---> Using cache
 ---> a85f650963c0
Step 18/18 : ENTRYPOINT ["/app/server"]
 ---> Using cache
 ---> 3c4e3917cad2
Successfully built 3c4e3917cad2
Successfully tagged go-template:latest
```

### Run the build

Take the latest build tag, `3c4e3917cad2` in the example above, and run a `docker run` command

```shell
docker run --env-file ./.env -p 50050:50050 3c4e3917cad2
```

This will run the image locally. You will be able to send an RPC to `grpc://127.0.0.1:50051`

and confirm the image is successfully working.
