schema "main" {}

table "races" {
  schema = schema.main

  column "id" {
    type = int
    null = false
  }
  column "slug" {
    type = text
    null = false
  }
  column "name" {
    type = text
    null = false
  }
  column "parent_race_id" {
    type = int
    null = true
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_race_parent" {
    columns     = [column.parent_race_id]
    ref_columns = [table.races.column.id]
    on_delete   = SET_NULL
  }

  index "idx_races_slug" {
    columns = [column.slug]
    unique  = true
  }
}

table "name_styles" {
  schema = schema.main

  column "id" {
    type = int
    null = false
  }
  column "race_id" {
    type = int
    null = false
  }
  column "slug" {
    type = text
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_style_race" {
    columns     = [column.race_id]
    ref_columns = [table.races.column.id]
    on_delete   = CASCADE
  }
}

table "name_patterns" {
  schema = schema.main

  column "id" {
    type = int
    null = false
  }
  column "race_id" {
    type = int
    null = false
  }
  column "style_id" {
    type = int
    null = true
  }
  column "order" {
    type = int
    null = false
  }
  column "component_type" {
    type = text
    null = false
  }
  column "required" {
    type    = int
    null    = false
    default = 1
  }
  column "max_count" {
    type    = int
    null    = false
    default = 1
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_pattern_race" {
    columns     = [column.race_id]
    ref_columns = [table.races.column.id]
    on_delete   = CASCADE
  }

  foreign_key "fk_pattern_style" {
    columns     = [column.style_id]
    ref_columns = [table.name_styles.column.id]
    on_delete   = CASCADE
  }

  index "idx_pattern_race" {
    columns = [column.race_id, column.order]
  }
}

table "name_components" {
  schema = schema.main

  column "id" {
    type = integer
    null = false
  }
  column "race_id" {
    type = int
    null = false
  }
  column "component_type" {
    type = text
    null = false
  }
  column "gender" {
    type = text
    null = true
  }
  column "value" {
    type = text
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_component_race" {
    columns     = [column.race_id]
    ref_columns = [table.races.column.id]
    on_delete   = CASCADE
  }

  index "idx_component_lookup" {
    columns = [column.race_id, column.component_type, column.gender]
  }
}

table "composite_parts" {
  schema = schema.main

  column "id" {
    type = integer
    null = false
  }
  column "race_id" {
    type = int
    null = false
  }
  column "position" {
    type = text
    null = false
  }
  column "category" {
    type = text
    null = false
  }
  column "value" {
    type = text
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_part_race" {
    columns     = [column.race_id]
    ref_columns = [table.races.column.id]
    on_delete   = CASCADE
  }
}
