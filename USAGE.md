<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
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

    ctx := context.Background()
    res, err := s.MatchmakerLobbies.LobbiesServiceFind(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MatchmakerFindLobbyOutput != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->