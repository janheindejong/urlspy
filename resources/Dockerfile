FROM python:3.10

ENV PYTHONDONTWRITEBYTECODE 1 \
    PYTHONUNBUFFERED 1

RUN apt-get update \
    && apt-get install curl -y \
    && curl -sSL https://install.python-poetry.org | python -

ENV PATH="/root/.local/bin:$PATH"

WORKDIR /usr/resources

COPY pyproject.toml poetry.lock ./

RUN poetry config virtualenvs.create false \
    && poetry install --no-dev --no-interaction --no-ansi

# Copy application source code
COPY ./app ./app
COPY entry-point.sh alembic.ini ./
COPY alembic alembic

# Modify entry-point permissions
RUN chmod a+x entry-point.sh

EXPOSE 8080

ENTRYPOINT ["./entry-point.sh"]

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8080"]