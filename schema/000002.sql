CREATE OR REPLACE FUNCTION fieldkit.fk_clustered_docs(source_id BIGINT)
RETURNS TABLE (
	"id" BIGINT,
	"input_id" INTEGER,
	"timestamp" timestamp,
	"time_difference" FLOAT,
	"temporal_cluster_id" BIGINT,
	"location" geometry,
	"spatial_cluster_id" integer
) AS
'
BEGIN
return query
WITH
source AS (
	SELECT
		d.id,
		d.input_id,
		d.timestamp,
		d.location
	FROM fieldkit.document d WHERE d.input_id IN (source_id) AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0
),
with_timestamp_differences AS (
	SELECT
		*,
			                              LAG(s.timestamp) OVER (PARTITION BY s.input_id ORDER BY s.input_id, s.timestamp) AS previous_timestamp,
		EXTRACT(epoch FROM (s.timestamp - LAG(s.timestamp) OVER (PARTITION BY s.input_id ORDER BY s.input_id, s.timestamp))) AS time_difference
	FROM source s
	ORDER BY s.input_id, s.timestamp
),
with_temporal_clustering AS (
	SELECT
		*,
		CASE WHEN s.time_difference > 600
			OR s.time_difference IS NULL THEN true
			ELSE NULL
		END AS new_temporal_cluster
	FROM with_timestamp_differences s
),
with_assigned_temporal_clustering AS (
	SELECT
		*,
		COUNT(new_temporal_cluster) OVER (
			PARTITION BY s.input_id
			ORDER BY s.input_id, s.timestamp
			ROWS UNBOUNDED PRECEDING
		) AS temporal_cluster_id
	FROM with_temporal_clustering s
),
spatial_clustering AS (
	SELECT
		s.id,
		s.input_id,
		s.timestamp,
		s.time_difference,
		s.temporal_cluster_id,
		s.location,
		ST_ClusterDBSCAN(ST_Transform(s.location, 2950), eps := 50, minpoints := 10) OVER (PARTITION BY s.input_id ORDER BY s.input_id, s.timestamp) AS spatial_cluster_id
	FROM with_assigned_temporal_clustering s
)
SELECT
	s.*
FROM spatial_clustering s;
END
' LANGUAGE plpgsql;

CREATE TABLE fieldkit.sources_temporal_clusters (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    cluster_id integer NOT NULL,
    updated_at timestamp NOT NULL,
    number_of_features integer NOT NULL,
    start_time timestamp NOT NULL,
    end_time timestamp NOT NULL,
    centroid geometry(POINT, 4326) NOT NULL,
    radius decimal NOT NULL,
    PRIMARY KEY (source_id, cluster_id)
);

CREATE TABLE fieldkit.sources_temporal_geometries (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    cluster_id integer NOT NULL,
    updated_at timestamp NOT NULL,
    geometry geometry(LINESTRING, 4326) NOT NULL,
    PRIMARY KEY (source_id, cluster_id)
);

CREATE TABLE fieldkit.sources_spatial_clusters (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    cluster_id integer NOT NULL,
    updated_at timestamp NOT NULL,
    number_of_features integer NOT NULL,
	start_time timestamp NOT NULL,
	end_time timestamp NOT NULL,
    centroid geometry(POINT, 4326) NOT NULL,
    radius decimal NOT NULL,
    PRIMARY KEY (source_id, cluster_id)
);