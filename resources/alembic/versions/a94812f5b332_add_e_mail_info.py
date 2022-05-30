"""Add e-mail info

Revision ID: a94812f5b332
Revises: b5518039a948
Create Date: 2022-05-30 17:39:43.626986

"""
import sqlalchemy as sa

from alembic import op

# revision identifiers, used by Alembic.
revision = "a94812f5b332"
down_revision = "b5518039a948"
branch_labels = None
depends_on = None


def upgrade():
    op.add_column("resource", sa.Column("email_address", sa.String))


def downgrade():
    op.drop_column("resource", "email_address")
