from fastapi import FastAPI
from fastapi import Body
from pydantic import BaseModel
from random import randrange

app = FastAPI()


posteos = [
    {"t": "titulo1", "c": "contenido", "id": 1},
    {"t": "titulo2", "c": "contenido2", "id": 2},
]


class Post(BaseModel):
    title: str
    content: str


@app.get("/")
async def root():
    return {"Message ": "hola mundillo"}


@app.get("/postear")
def get_mensajes():
    return {"data": posteos}


@app.post("/postear")
def postear(posts: Post):
    posteo = posts.dict()
    posteo["id"] = randrange(0, 10000)
    posteos.append(posteo)
    return {"data": posteo}


@app.put("/postear/{id}")
def update(posts: Post):
    pass


@app.get("/postear/{id}")
def get_post_id(id):
    post = post_id(int(id))
    return {"Detalles": post}


def post_id(id):
    for p in posteos:
        if p["id"] == id:
            return p


@app.delete("/postear/{id}")
def delete_post(id: int):
    for p in posteos:
        if p["id"] == id:
            posteos.pop(p)
