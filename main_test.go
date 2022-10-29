package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func Test_acceptanceTest(t *testing.T) {
	type args struct {
		route string
		body  string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		want     string
	}{
		{
			name: "happy path",
			args: args{
				route: "/v1/maximize",
				body:  "[\n    {\n        \"request_id\": \"bookata_XY123\",\n        \"check_in\": \"2020-01-01\",\n        \"nights\": 5,\n        \"selling_rate\": 200,\n        \"margin\": 20\n    },\n    {\n        \"request_id\": \"kayete_PP234\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 156,\n        \"margin\": 5\n    },\n    {\n        \"request_id\": \"atropote_AA930\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 150,\n        \"margin\": 6\n    },\n    {\n        \"request_id\": \"acme_AAAAA\",\n        \"check_in\": \"2020-01-10\",\n        \"nights\": 4,\n        \"selling_rate\": 160,\n        \"margin\": 30\n    }\n]",
			},
			wantCode: fiber.StatusOK,
			want:     "{\"request_ids\":[\"bookata_XY123\",\"acme_AAAAA\"],\"total_profit\":88,\"avg_night\":10,\"min_night\":8,\"max_night\":12}",
		},
		{
			name: "invalid date",
			args: args{
				route: "/v1/maximize",
				body:  "[\n    {\n        \"request_id\": \"bookata_XY123\",\n        \"check_in\": \"2020-01-1\",\n        \"nights\": 5,\n        \"selling_rate\": 200,\n        \"margin\": 20\n    },\n    {\n        \"request_id\": \"kayete_PP234\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 156,\n        \"margin\": 5\n    },\n    {\n        \"request_id\": \"atropote_AA930\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 150,\n        \"margin\": 6\n    },\n    {\n        \"request_id\": \"acme_AAAAA\",\n        \"check_in\": \"2020-01-10\",\n        \"nights\": 4,\n        \"selling_rate\": 160,\n        \"margin\": 30\n    }\n]",
			},
			wantCode: fiber.StatusBadRequest,
			want:     "{\"error\":true,\"msg\":\"There's a problem with the input data Key: 'BookingsRequestDTO.Checkin' Error:Field validation for 'Checkin' failed on the 'len' tag\"}",
		},
		{
			name: "invalid nights",
			args: args{
				route: "/v1/maximize",
				body:  "[\n    {\n        \"request_id\": \"bookata_XY123\",\n        \"check_in\": \"2020-01-01\",\n        \"nights\": -5,\n        \"selling_rate\": 200,\n        \"margin\": 20\n    },\n    {\n        \"request_id\": \"kayete_PP234\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 156,\n        \"margin\": 5\n    },\n    {\n        \"request_id\": \"atropote_AA930\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 150,\n        \"margin\": 6\n    },\n    {\n        \"request_id\": \"acme_AAAAA\",\n        \"check_in\": \"2020-01-10\",\n        \"nights\": 4,\n        \"selling_rate\": 160,\n        \"margin\": 30\n    }\n]",
			},
			wantCode: fiber.StatusBadRequest,
			want:     "{\"error\":true,\"msg\":\"There's a problem with the input data Key: 'BookingsRequestDTO.Nights' Error:Field validation for 'Nights' failed on the 'min' tag\"}",
		},
		{
			name: "wrong path",
			args: args{
				route: "/maximize",
				body:  "[\n    {\n        \"request_id\": \"bookata_XY123\",\n        \"check_in\": \"2020-01-01\",\n        \"nights\": 5,\n        \"selling_rate\": 200,\n        \"margin\": 20\n    },\n    {\n        \"request_id\": \"kayete_PP234\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 156,\n        \"margin\": 5\n    },\n    {\n        \"request_id\": \"atropote_AA930\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 150,\n        \"margin\": 6\n    },\n    {\n        \"request_id\": \"acme_AAAAA\",\n        \"check_in\": \"2020-01-10\",\n        \"nights\": 4,\n        \"selling_rate\": 160,\n        \"margin\": 30\n    }\n]",
			},
			wantCode: fiber.StatusNotFound,
			want:     "{\"error\":true,\"msg\":\"sorry, endpoint is not found\"}",
		},
		{
			name: "happy path stats",
			args: args{
				route: "/v1/stats",
				body:  "[\n    {\n        \"request_id\": \"bookata_XY123\",\n        \"check_in\": \"2020-01-01\",\n        \"nights\": 5,\n        \"selling_rate\": 200,\n        \"margin\": 20\n    },\n    {\n        \"request_id\": \"kayete_PP234\",\n        \"check_in\": \"2020-01-04\",\n        \"nights\": 4,\n        \"selling_rate\": 156,\n        \"margin\": 22\n    }\n]",
			},
			wantCode: fiber.StatusOK,
			want:     "{\"avg_night\":8.29,\"min_night\":8,\"max_night\":8.58}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			SetupRoutes(app)
			req := httptest.NewRequest("POST", tt.args.route, strings.NewReader(tt.args.body))
			req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
			req.Header.Add("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.Equal(t, tt.wantCode, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want, string(body))
		})
	}
}
