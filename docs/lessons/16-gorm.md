# 16 · GORM — ORM for Go

Coming from Java: GORM ≈ a lightweight JPA/Hibernate — struct tags instead of annotations, explicit where Hibernate is "magic". Know both its power and its footguns.

## Setup & Models

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

type User struct {
    ID        uint           `gorm:"primaryKey"`            // or embed gorm.Model
    Name      string         `gorm:"size:100;not null;index"`
    Email     string         `gorm:"uniqueIndex"`
    Orders    []Order        // has-many
    CreatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`                  // soft delete
}

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Warn), // log slow SQL
})
db.AutoMigrate(&User{}, &Order{})   // DDL sync — dev only, use migrations in prod
```

## CRUD & queries

```go
// Create
db.Create(&User{Name: "Alice", Email: "a@b.c"})

// Read — First vs Find vs Take (First adds ORDER BY id LIMIT 1)
var u User
result := db.First(&u, "email = ?", "a@b.c")   // NEVER fmt.Sprintf into SQL
if errors.Is(result.Error, gorm.ErrRecordNotFound) { /* 404 */ }

// Update
db.Model(&u).Update("name", "Alicia")                       // single field
db.Model(&u).Updates(User{Name: "X"})                      // skips zero fields!
db.Model(&u).Updates(map[string]any{"name": ""})           // map = zero values too

// Delete (soft if DeletedAt exists)
db.Delete(&u)
db.Unscoped().Delete(&u)                                   // hard delete
```

## Associations — the N+1 trap

```go
type Order struct {
    ID     uint
    UserID uint
    Items  []Item
}

// BAD: N+1 — one query for users, then N queries for their orders
var users []User
db.Find(&users)
for _, u := range users { db.Find(&u.Orders) }   // N queries!

// GOOD: Preload (like JPA fetch join / eager loading)
db.Preload("Orders").Preload("Orders.Items").Find(&users)

// Selective preload
db.Preload("Orders", "status = ?", "shipped").Find(&users)
```

## Transactions

```go
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&order).Error; err != nil {
        return err                 // any error → rollback
    }
    if err := tx.Model(&account).Update("balance", gorm.Expr("balance - ?", amt)).Error; err != nil {
        return err
    }
    return nil                     // nil → commit
})
// Never swallow tx errors; the closure pattern handles rollback for you
```

## Pitfalls — the interview checklist

!!! warning "Zero-value blind spot"
    `Updates(struct)` skips zero fields (`0`, `""`, `false`, nil time). Setting a field **to** zero requires `map[string]any` or `Select("*")`. #1 GORM bug source.

!!! warning "Scopes"
    `db.Where(...)` returns a *new* chain — reusing a `tx := db.Where(...)` accumulates conditions across calls. Use `Session(&gorm.Session{})` or rebuild per request.

!!! warning "ErrRecordNotFound"
    `First` returns `gorm.ErrRecordNotFound`; `Find` on slices does **not** (empty slice, nil error). Know the difference — it's a real bug pattern.

!!! warning "Skip hooks unintentionally"
    `db.Model(&u).UpdateColumn(...)` skips `BeforeUpdate` hooks and `UpdatedAt`. Sometimes what you want; often a surprise.

!!! warning "Connection pool"
    `sqlDB, _ := db.DB(); sqlDB.SetMaxOpenConns(25)` — GORM defaults can starve under load. Tune `MaxOpenConns`, `MaxIdleConns`, `ConnMaxLifetime`.

## Alternatives worth naming in an interview

- **`sqlx`** — thin SQL + struct scanning, no magic.
- **`sqlc`** — write SQL, codegen type-safe Go. Many teams prefer this; GORM critics cite "impedance mismatch hidden until prod".
- **`ent`** (Facebook) — entity codegen, type-safe graph queries.
- **`pgx`** — if Postgres-only and want raw performance.

**Interview line:** "I'd default to `sqlc`/`pgx` for new services and use GORM where the domain genuinely benefits from an ORM — always with `DryRun`/SQL logging in tests to catch N+1s."

<quiz>
`db.Model(&u).Updates(User{Name: "", Age: 0})` — what gets updated?
- [ ] Both fields set to empty/0
- [x] Nothing — struct-based Updates skips zero-value fields (use a map)
- [ ] Compile error
- [ ] Only Name
</quiz>

<quiz>
How do you eager-load associations in GORM?
- [ ] `@Fetch(EAGER)` annotation
- [x] `.Preload("Orders")`
- [ ] `.Join()` always
- [ ] Automatic by default
</quiz>

<quiz>
`db.Find(&users)` finds nothing — what's `err`?
- [ ] `gorm.ErrRecordNotFound`
- [x] `nil` (empty slice is not an error — unlike `First`)
- [ ] `sql.ErrNoRows`
- [ ] panic
</quiz>
