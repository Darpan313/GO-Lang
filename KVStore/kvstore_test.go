package kvstore

import "testing"

func TestNestedTransactions(t *testing.T) {
	db := NewDB()

	// Initial State
	db.Set("a", "1")
	db.Set("b", "2")

	// START TX 1
	db.Begin()
	db.Set("a", "10")	// Override
	db.Delete("b")		// Tombstone
	db.Set("c", "30")	// New

	assertGet(t, db, "a", "10", true)
	assertGet(t, db, "b", "", false)
	assertGet(t, db, "c", "30", true)

	// START TX 2 (Nested)
	db.Begin()
	db.Set("a", "100")	// Override parent tx
	db.Set("b", "200")	// Resurrect deleted key

	assertGet(t, db, "a", "100", true)
	assertGet(t, db, "b", "200", true)

	// ROLLBACK TX 2
	err := db.Rollback()
	if err != nil {
		t.Fatalf("unexpected error on rollback: %v", err)
	}

	// Verify we are back to TX 1 state
	assertGet(t, db, "a", "10", true)
	assertGet(t, db, "b", "", false)
	assertGet(t, db, "c", "30", true)

	// START TX 3 (Nested)
	db.Begin()
	db.Set("d", "40")

	// COMMIT TX 3 into TX 1
	err = db.Commit()
	if err != nil {
		t.Fatalf("unexpected error on commit: %v", err)
	}

	// COMMIT TX 1 into base store
	err = db.Commit()
	if err != nil {
		t.Fatalf("unexpected error on commit: %v", err)
	}

	// Verify final base store state
	assertGet(t, db, "a", "10", true)
	assertGet(t, db, "b", "", false)
	assertGet(t, db, "c", "30", true)
	assertGet(t, db, "d", "40", true)

	// Verify empty transaction stack error
	err = db.Commit()
	if err != ErrNoTransaction {
		t.Fatalf("expected ErrNoTransaction, got %v", err)
	}
}

func assertGet(t *testing.T, db *DB, key, expectedValue string, expectedExists bool) {
	t.Helper()
	val, exists := db.Get(key)

	if exists != expectedExists {
		t.Errorf("key '%s': expected exists=%v, got %v", key, expectedExists, exists)
	}

	if exists && val != expectedValue {
		t.Errorf("key '%s': expected value='%s', got '%s'", key, expectedValue, val)
	}
}