FROM python:3.10

ENV PYTHONDONTWRITEBYTECODE 1 \
    PYTHONUNBUFFERED 1

# Install Poetry
RUN apt-get update \
    && apt-get install curl -y \
    && curl -sSL https://install.python-poetry.org | python -
ENV PATH="/root/.local/bin:$PATH"

# Install Python dependencies
WORKDIR /workspace
COPY pyproject.toml poetry.lock ./
RUN poetry config virtualenvs.create false \
    && poetry install --no-dev --no-interaction --no-ansi

# Copy application source code
COPY . .

# Expose application
EXPOSE 80

# Run app
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "80"]
