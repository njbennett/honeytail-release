# Honeytail BOSH Release

Welcome to the highly experimental Honeytail BOSH release

# WARNINGS
This release runs Honeytail as the root user.
It is probably unsafe, especially for production systems.

It currently depends on knowing exactly which log files you want it to run.
Log files are not usually part of a release's public interface,
so this is likely to silently break.

The structure is likely to change a *lot*
and/or go unmaintained for long periods of time.
So just, y'know, keep that in mind
before deploying somewhere you care about.

# To Use
To build the release,
you'll need to download the honeytail binary yourself
and run `bosh add-blob` to make it available to the director.
(Probably-- I haven't actually tested that this works
on anything other than my dev machine yet.)

Apply the `add-honeytail.yml` opsfile to cf-deployment
It was tested with cf-deployment 4.3.0
