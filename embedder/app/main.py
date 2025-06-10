from fastapi import FastAPI
from app.schema import embed_text
from app.model import EmbedRequest, EmbedResponse

app = FastAPI()


@app.post("/embed", response_model=EmbedResponse)
def embed(request: EmbedRequest) -> EmbedResponse:
    """
    Endpoint to embed text and return the vector.
    """
    vector = embed_text(request.text)
    return {"vector": vector}
