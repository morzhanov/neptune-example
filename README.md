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

