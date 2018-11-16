# Honeytail BOSH Release

Welcome to the highly experimental Honeytail BOSH release.

Let me know if you're using it! I can be found at nbennett@pivotal.io

# WARNINGS
This release uses the unsafe `unrestricted_volumes` feature of BPM
to give the honeytail job access to all logs on the job it is deployed with.

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

Apply the `add-honeytail.yml` opsfile to cf-deployment
It was tested with cf-deployment 4.3.0
