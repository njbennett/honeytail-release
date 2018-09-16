# Honeytail BOSH Release

Welcome to the highly experimental Honeytail BOSH release

# WARNINGS
This release runs Honeytail as the root user.
It is probably unsafe, especially for production systems.

It currently depends on knowing exactly which log files you want it to run.
Log files are not usually part of a release's public interface,
so this is likely to silently break.

# To Use
Apply the `add-honeytail.yml` opsfile to cf-deployment
It was tested with cf-deployment 4.3.0
