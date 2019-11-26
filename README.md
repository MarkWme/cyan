# Cyan API Server and Applications

This repository contains a simple API server that returns a JSON payload when a GET is performed against the /api/getVersion endpoint. The payload is simply application version information. There are three client applications, written in Go, Node.js and Python, that hit the endpoint of the server once per second and output the payload returned, or "Error" if there is a problem connecting.

These applications can be used for simple testing of a blue / green / canary deployment scenario. By deploying and running the client applications, they will continually poll the endpoint. You can modify the version information returned by the endpoint to simulate deployment of a new version of the server application and then experiment with different ways of deploying, testing and releasing the new application.

The apps are designed to be simple so that focused is maintained on experimentation with deployment scenarios and not on the apps themselves.