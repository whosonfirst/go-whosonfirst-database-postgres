package create

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var iterator_uri string

var database_uri string

var all bool
var ancestors bool
var concordances bool
var geojson bool
var spelunker bool
var geometries bool
var names bool
var rtree bool
var properties bool
var search bool
var spr bool
var supersedes bool

var spatial_tables bool
var spelunker_tables bool

var alt_files bool
var strict_alt_files bool

var index_alt multi.MultiString

var index_relations bool
var relations_uri string

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("index")

	fs.StringVar(&database_uri, "database-uri", "", "...")

	fs.BoolVar(&all, "all", false, "Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)")
	fs.BoolVar(&ancestors, "ancestors", false, "Index the 'ancestors' tables")
	fs.BoolVar(&concordances, "concordances", false, "Index the 'concordances' tables")
	fs.BoolVar(&geojson, "geojson", false, "Index the 'geojson' table")
	fs.BoolVar(&spelunker, "spelunker", false, "Index the 'spelunker' table")
	fs.BoolVar(&geometries, "geometries", false, "Index the 'geometries' table (requires that libspatialite already be installed)")
	fs.BoolVar(&names, "names", false, "Index the 'names' table")
	fs.BoolVar(&rtree, "rtree", false, "Index the 'rtree' table")
	fs.BoolVar(&properties, "properties", false, "Index the 'properties' table")
	fs.BoolVar(&search, "search", false, "Index the 'search' table (using SQLite FTS4 full-text indexer)")
	fs.BoolVar(&spr, "spr", false, "Index the 'spr' table")
	fs.BoolVar(&supersedes, "supersedes", false, "Index the 'supersedes' table")

	fs.BoolVar(&spatial_tables, "spatial-tables", false, "If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.")
	fs.BoolVar(&spelunker_tables, "spelunker-tables", false, "If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages")

	fs.BoolVar(&alt_files, "index-alt-files", false, "Index alt geometries. This flag is deprecated, please use -index-alt=TABLE,TABLE,etc. instead. To index alt geometries in all the applicable tables use -index-alt=*")
	fs.Var(&index_alt, "index-alt", "Zero or more table names where alt geometry files should be indexed.")

	fs.BoolVar(&strict_alt_files, "strict-alt-files", true, "Be strict when indexing alt geometries")

	fs.BoolVar(&index_relations, "index-relations", false, "Index the records related to a feature, specifically wof:belongsto, wof:depicts and wof:involves. Alt files for relations are not indexed at this time.")
	fs.StringVar(&relations_uri, "index-relations-reader-uri", "", "A valid go-reader.Reader URI from which to read data for a relations candidate.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging")
	return fs
}
