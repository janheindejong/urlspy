"""Create resource table

Revision ID: b5518039a948
Revises: 
Create Date: 2022-05-13 08:17:38.030277

"""
import sqlalchemy as sa

from alembic import op

# revision identifiers, used by Alembic.
revision = "b5518039a948"
down_revision = None
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        "resource",
        sa.Column("id", sa.Integer, primary_key=True),
        sa.Column("url", sa.String, nullable=False),
    )


def downgrade():
    op.drop_table("resource")
