# neptune-example

AWS Neptune Go application example.

## Graph Data Schema

### Vertexes

| Vertex   |      Description      |
|----------|:---------------------:|
| Reader   |  User Reader          |
| Book     |  A Book               |
| Author   |  User Writer          |

### Edges

| Edge     |  Source  |   Target |
|----------|:--------:|:--------:|
| read     |  Reader  | Book     |
| authored |  Author  | Book     |

### Example

<img src="https://i.ibb.co/745F7PZ/IMG-0154.png" alt="example graph"/>

// TODO: we will use wss endpoint to gremlin
// TODO: neptune should be publicitly available via wss endpoint for this example
// TODO: more info: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin.html

## How it works

Simple application which uses AWS Neptune database as persistent storage. Authors create Books, Readers read Books.

Application uses Gremlin language which is supported by <a href="https://github.com/northwesternmutual/grammes">grammes</a> package. For more info about Gremlin language refer to the <a href="http://tinkerpop.apache.org/docs/current/reference/#_tinkerpop_documentation">TinkerPop</a> docs.

For this application we are using AWS Neptune publicity available WebSockets endpoint. For more information about setup refer to the <a href="https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin.html">AWS Docs</a>.

## Structure

- `/cmd` - application setup
- `/config` - .env file with environment variables
- `/internal`
    - `/config` - config module based on viper package
    - `/logger` - application logger, creates file transport (for filebeat) and console transport
    - `/db` - Neptune database connection
    - `/service` - main app service
    