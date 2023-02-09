# Rivet Go SDK

The backbone of Multiplayer Experiences.  [Simple APIs for serverless lobbies, matchmaking, CDN, and other features](https://docs.rivet.gg/general/introduction) managed directly within an easy to use dashboard. Launch your multiplayer game on any platform with the tools you already use.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/rivet-go-sdk
```
<!-- End SDK Installation -->

## Authentication

Signup for [access]([https://www.leapml.dev/signup](https://hub.rivet.gg/developer/dashboard)) to Rivet to use the API. 

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/rivet-go-sdk"
    "github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/operations"
)

func main() {
    s := rivet.New()
    
    req := operations.LobbiesServiceFindRequest{
        Security: operations.LobbiesServiceFindSecurity{
            BearerAuth: &shared.SchemeBearerAuth{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
        Headers: operations.LobbiesServiceFindHeaders{
            Origin: "unde",
        },
        Request: operations.LobbiesServiceFindRequestBody{
            Captcha: "deserunt",
            GameModes: []string{
                "nulla",
                "id",
                "vero",
            },
            PreventAutoCreateLobby: false,
            Regions: []string{
                "nihil",
                "fuga",
                "facilis",
                "eum",
            },
        },
    }
    
    res, err := s.MatchmakerLobbies.LobbiesServiceFind(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MatchmakerFindLobbyOutput != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### MatchmakerLobbies

* `LobbiesServiceFind` - Finds a lobby based on the given criteria.
If a lobby is not found and `prevent_auto_create_lobby` is `true`, 
a new lobby will be created.

* `LobbiesServiceJoin` - Joins a specific lobby.
This request will use the direct player count configured for the
lobby group.

* `LobbiesServiceList` - Lists all open lobbies.
* `LobbiesServiceSetClosed` - If `is_closed` is `true`, players will be prevented from joining the lobby.
Does not shutdown the lobby.


### MatchmakerRegions

* `RegionsServiceList` - Returns a list of regions available to this namespace.

Regions are sorted by most optimal to least optimal. The player's IP address
is used to calculate the regions' optimality.

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
