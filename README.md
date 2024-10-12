# pcr
Phone configuration registry. This is a service that dynamically renders the appropriate configuration file
for a given hard-phone based on its User-Agent. It relies on the fact that most phones state their model and
MAC address in their user agent.

This makes it possible to customize the configuration file provided based on the phone's identity, with the main
benefit being that credentials can be provided for each phone that requests its configuration.

## Supported models
All User-Agents are different. This is a list of current models which are supported:

- Avaya J100 series.

## Setup
This application is not meant to be deployed as a standalone service. Production deployments should use a reverse proxy
that provides TLS support. It is also recommended to secure this service through something like HTTP basic
authentication which should be supported by most hard-phones. Consider other security measures like rate-limiting
as well. If an attacker gets access to this service he can obtain credentials for your PBX if he manages
to spoof the user-agent for one the phones in the database.
