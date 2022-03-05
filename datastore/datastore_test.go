package datastore

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/t-bonatti/license-manager/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("DataStore", func() {
	var datastore *dataStoreImpl
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{}) // open gorm db
		Expect(err).ShouldNot(HaveOccurred())

		datastore = &dataStoreImpl{db: gdb}
	})
	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("get", func() {

		It("when is empty", func() {
			const sqlSelectAll = `SELECT * FROM "licenses" WHERE id = $1 AND version = $2`
			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
				WithArgs("1", "v1").
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := datastore.Get("1", "v1")
			Expect(err).To(MatchError(gorm.ErrRecordNotFound))
		})

		It("when has data", func() {
			licence := &model.License{
				ID:        "1",
				Version:   "v10",
				CreatedAt: time.Now(),
				Info:      types.JSONText(`{"company": "xyz"}`),
			}

			rows := sqlmock.
				NewRows([]string{"id", "version", "created_at", "info"}).
				AddRow(licence.ID, licence.Version, licence.CreatedAt, licence.Info)

			const sqlSelectAll = `SELECT * FROM "licenses" WHERE id = $1 AND version = $2`
			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
				WithArgs("1", "v10").
				WillReturnRows(rows)

			dbLicense, err := datastore.Get("1", "v10")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dbLicense.ID).To(Equal(licence.ID))
			Expect(dbLicense.Version).To(Equal(licence.Version))
			Expect(dbLicense.CreatedAt).To(Equal(licence.CreatedAt))
			Expect(dbLicense.Info).To(Equal(licence.Info))
		})

	})
})
