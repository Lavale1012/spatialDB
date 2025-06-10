from sentence_transformers import SentenceTransformer

# Load the model once when the app starts
model = SentenceTransformer("all-MiniLM-L6-v2")

# Function that takes in text and returns an embedding (vector)


def embed_text(text: str) -> list[float]:
    embedding = model.encode(text)
    return embedding.tolist()
