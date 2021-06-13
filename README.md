
[![Go](https://github.com/ayushjha-fm/tmai.server.mock/actions/workflows/go.yml/badge.svg)](https://github.com/ayushjha-fm/tmai.server.mock/actions/workflows/go.yml)
[![Publish Docker image](https://github.com/ayushjha-fm/tmai.server.mock/actions/workflows/docker.yml/badge.svg)](https://github.com/ayushjha-fm/tmai.server.mock/actions/workflows/docker.yml)

```

╔════╗╔═╗╔═╗╔═══╗╔══╗    ╔═╗╔═╗        ╔╗      ╔═══╗
║╔╗╔╗║║║╚╝║║║╔═╗║╚╣╠╝    ║║╚╝║║        ║║      ║╔═╗║
╚╝║║╚╝║╔╗╔╗║║║ ║║ ║║     ║╔╗╔╗║╔══╗╔══╗║║╔╗    ║╚══╗╔══╗╔═╗╔╗╔╗╔══╗╔═╗
  ║║  ║║║║║║║╚═╝║ ║║     ║║║║║║║╔╗║║╔═╝║╚╝╝    ╚══╗║║╔╗║║╔╝║╚╝║║╔╗║║╔╝
 ╔╝╚╗ ║║║║║║║╔═╗║╔╣╠╗    ║║║║║║║╚╝║║╚═╗║╔╗╗    ║╚═╝║║║═╣║║ ╚╗╔╝║║═╣║║
 ╚══╝ ╚╝╚╝╚╝╚╝ ╚╝╚══╝    ╚╝╚╝╚╝╚══╝╚══╝╚╝╚╝    ╚═══╝╚══╝╚╝  ╚╝ ╚══╝╚╝


```


Table of Contents
=================
* [Running the mock server](#running-the-mock-server)
* [Using the Server](#using-the-server)
* [Some example cURL commands](#some-example-curl-commands)
* [Some example fetch commands](#some-example-fetch-commands)
* [How does it work](#how-does-it-work)
* [Features](#features)
* [Structure of api.json](#structure-of-apijson)
* [Structure of Endpoint](#structure-of-endpoint)
* [Customize messages type](#customize-messages-type)
* [Customize suggestions type](#customize-suggestions-type)
* [License](#license)


### Running the mock server

```bash
## Pull the image separately so you always have the latest version
docker pull ayushjhafm/tmai.server.mock:latest
docker run -p 3000:3000 ayushjhafm/tmai.server.mock
```

### Using the Server

The above command should start a server on port 3000, so now you can go to <http://localhost:3000/messages>
for messages response, and <http://localhost:3000/trivia> for trivia response.

#### Some example cURL commands

Input
```bash
curl -H"Suggestions: Explore-Questions" -H"Messages: text, article" http://localhost:3000/message
```

Output
```json
{
  "query": "",
  "meta": null,
  "data": [
    {
      "type": "text",
      "meta": {
        "...": "...contains meta info that may ne useful in the future"
      },
      "content": "Benedicis cor adest falli tu ergone bestiae id eruuntur, ob da cor passionis da duabus manifestetur audi ita da, iugo cor o agnosco opus condit se sumus."
    },
    {
      "type": "text",
      "meta": {
        "...": "...contains meta info that may ne useful in the future"
      },
      "title": "Plus esau iam fias quaeso se vestibus, iniuste agenti ait aliquid via evidentius.",
      "content": "Me places omne os a capiendum mala ardentius. Cor genus hac magnifico iamque recti sed delectari, agam cervicem. Lux respuo ita hi isto os contra periculis non agnoscerent te ita nonnullius servis dum da. Fuero dignationem poenaliter divino praeciderim, petimus mansuefecisti dicat quaeso, contra cor sua rerum gaudeam. Plenariam voluptaria facultas, hi tutor ore da, necessarium ei. Temporis beares erigo proprios temptatum me terram de diutius diu. Tale eius medice in seu ipsa spe re aliquantum inter nam hi da places te a ob quotiens possint removeri. Es o re, vivam saucio bonorumque eosdem fit locutus cor me. Ego re recti ut singillatim. Duxi gaudium ita e muta consulerem vim me dulcedine prout vestibus causa cui sat. Propria oculi sed tale id a re odor singula quaeritur hi, procedens rogo. Foeda pro formosa, a mea. Me ingentibus re nam plenariam regem manet tetigisti suavitatem post gloriatur te dabis cor extra mortalitatis da agam transisse, ipsae te. Aut te singillatim horum at ore tibi infelix talia nam an edunt. Ante vi sanctuarium mediator me fama diligit agro re re. Agam nisi inferiora turpitudines, relinquentes verae. Pro qui usurpant lugens oblivionem an ore suis credimus spe cor hi extraneus nota ob excellentiam fuimus. Flammam iniquitatis primatum nomine tu intime eloquentes imaginatur intime periculis se dum aer, colligo te manducat soni da ob artibus. Hi cuius cura inconsummatus da. Regem re me sapores miracula corporales strepitu florum perdiderat locus absurdissimum denuo e ea se aer ita. Deum conor o nolo alimenta aut. Caeli ungentorum fac cuncta remota re, faciunt ita temptatione alteram latet admittantur sero at. Vestra familiaritate saucium flabiles dixi cibum ne re se. Voluptaria id agam te os nam a. O tu potui, da cervicem inconsummatus me compagem ita ob meo prodeat receptaculis.",
      "pub_date": "2020-12-12",
      "source_url": "times.com",
      "author": {
        "name": "Habes nam.",
        "twitter_username": "Divino."
      }
    }
  ],
  "suggestions": [
    {
      "type": "explore_questions"
    }
  ]
}
```

Input
```bash
curl -H"Suggestions: Suggest-Questions" -H"Messages: infographic" -d '{"query": "are you alive?"}' http://localhost:3000/message
```

Output
```json
{
  "query": "are you alive?",
  "meta": null,
  "data": [
    {
      "type": "infographics",
      "meta": {
        "...": "...contains meta info that may ne useful in the future"
      },
      "title": "Title of infographic message",
      "content": [
        {
          "name": "New York",
          "data": "URL TO IMAGE OR BASE64",
          "datatype": "URL / BASE64; specify the type of data"
        },
        {
          "name": "Gorgia",
          "data": "URL TO IMAGE OR BASE64",
          "datatype": "URL / BASE64; specify the type of data"
        },
        {
          "name": "Boston",
          "data": "URL TO IMAGE OR BASE64",
          "datatype": "URL / BASE64; specify the type of data"
        }
      ]
    }
  ],
  "suggestions": [
    {
      "type": "suggested_questions",
      "questions": [
        {
          "text": "<TEXT FOR QUESTION>"
        },
        {
          "text": "<TEXT FOR QUESTION>"
        },
        {
          "text": "<TEXT FOR QUESTION>"
        }
      ]
    }
  ]
}
```

Input
```bash
curl -H"Suggestions: No-Suggest" -H"Messages: suggested-topics" http://localhost:3000/message
```

Output
```json
{
  "query": "are you alive?",
  "meta": null,
  "data": [
    {
      "type": "text",
      "meta": {
        "...": "...contains meta info that may ne useful in the future"
      },
      "content": [
        {
          "text": "Aut adprobamus mei."
        },
        {
          "text": "Elapsum me."
        },
        {
          "text": "Solitis vitaliter fulgeat."
        }
      ]
    }
  ],
  "suggestions": null
}
```


#### Some example fetch commands

```javasript
await fetch("http://localhost:3000/message", {
    "credentials": "include",
    "headers": {
        "Messages": "text, article",
        "Suggestions": "Explore-Stabby"
    },
    "body": "{\n\"query\": \"hello, are you alive?\"\n}",
    "method": "GET",
});
```

### How does it work

The server serves JSON files in the `tmai` directory, as per the rules
described in `tmai/api.json`. This server is slightly more fancy than a simple JSON server.
The features for this mock server are described below.

#### Features
* Recieve and transmit the "query" from `request->response` for all the `message` type responses.
* Manage state of suggestions based on availability of `CONVERSATION_TOKEN`
* Read and send JSON responses by reading from JSON files with same name as the endpoint path.
  For example, the `/text` endpoint will read from `tmai/text.json` file.
* Customize the type of suggestion you are expecting from a message response.

### Structure of api.json

`api.json`

```javascript
{
    "endpoints": [{
            "path": "/text",
            "status": 200,
            "type": "message",
            "methods": [
                "GET",
                "POST"
            ]
    }],
    "port": 3000
}
```

#### Structure of Endpoint

| Field Name | Definition                                               |
|------------|----------------------------------------------------------|
| path       | The path for which this endpoint will serve a file       |
| status     | The status returned in response                          |
| type       | There are two types of responses, "message" and "trivia" |
| methods    | Which HTTP Methods should the endpoints support          |

### Customize messages type
By using `Messages` Header in the request, users can control the type of message they get.
There are four types of suggestions,

| Messages                      | Header Value     |
|-------------------------------|------------------|
| Article type message          | article          |
| Text Type Message             | text             |
| Infographic message           | infographic      |
| Suggested Topics gamification | suggested-topics |

By sending one or more of these as request, you can customize the type of messages you want.
For example, `Messages: text, article` will send both text and article messages with the
response in the specified order.


### Customize suggestions type

By using `Suggestions` Header in the request, users can control the type of suggestion they get.
There are four types of suggestions,

| Suggestion                            | Header Value      |
|---------------------------------------|-------------------|
| Explore Stabby Suggestion             | Explore-Stabby    |
| Explore popular quesitons with stabby | Explore-Questions |
| Get Suggested Questions               | Suggest-Questions |
| No Suggestions necessary              | No-Suggest        |

By sending one or more of these as request, you can customize the type of suggestions you want.
For example, `Suggestions: Explore-Stabby, Suggest-Questions` will send both explore stabby and
suggest-questions suggestions with the response in the specified order.

### License

MIT ✨
