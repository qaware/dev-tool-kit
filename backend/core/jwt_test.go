package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeJwt(t *testing.T) {
	// HS256
	valid, payload, err := DecodeJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.TfVxPqlJ2xhpB5mKN6fVvb51Z_mIPbdNw2D_1b_tN0w", "mysecret")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// HS384
	valid, payload, err = DecodeJwt("eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.X7Ecv7G0S8oK_DQG5zbRHJDT_r8qYzm2voJQtsB4NYnhuadHbJFa-09YTOTmKunB", "mysecret")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// HS512
	valid, payload, err = DecodeJwt("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.fx6w2G9MX2yjOIbvkcz-KisyjQ9gdKTCr4Mds_q1460bq_8PQMYZdhmmV0UIXDN-2bf0V4ALA6iNA1_qg0hdxA", "mysecret")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// RS256
	valid, payload, err = DecodeJwt("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.CZCZBy_YYP6xotUHZJzaWmo5B_d5FhjBAnNQvMfEg_EZm1G2P9hdfUkEUGl-Rdc2nvnrxhP6R4Hy4FPYRdsmheivjQVLUlh1rPN_o7bgNxZXRxUQLgGmnfBZJmmjZktT9w7O0NOqG6XZON584RddDDk2wrHb5t2RmlGjQeWP7zgss5lfIXZF_oSfXflUiMNAeAm-LQ6YL0s6xqa5dbCsd047uCHvIucGJOXUZl8lG2iy5NCViaC3DWlfrPk0C3VBv-kdfY0LY8xYiL6kroY4sdVvWqcHbH9yfZDDdYCk-xR5Q2U1f64ieVhxoI_0Vt0xTrcY1YNYKf0NFF5zYZJRKQ", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv\nvkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc\naT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy\ntvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0\ne+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb\nV6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9\nMwIDAQAB\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// RS256 with JWK
	valid, payload, err = DecodeJwt("eyJraWQiOiI1Y2Q2MWQ0NSIsIn\n    R5cCI6IkpXVCIsImFsZyI6IlJTMjU2In0.eyJuYW1lI\n  joiTWF4IE11c3Rlcm1hbm4iLCJpZCI6IjEyMzQ1Njc4OTAiLCJpYXQiOjE2MDY0Mjg1MDV9.U9WUtkDnNHe3SXGm1xJr8KfSZ9TvvnqmC3_r1akN5vHL2q73ZDf1ECjxNKj-Q6Oq4LJtJo9Hff0sbMsJXtLIfl0tvW4rB9gmdifu4KRj7QxfiCdUxdidaliYZXN_gCV10AIk8IvG7z_ITmbgd4-2hR4N-nqiSIsSfKmNcAyppvclXULFcw2_KqgefB1HQNasLiexH109YRWGTGIcSOS7OKEmjIeP-3Wd5kZJULJBWjryOJP6g-kNsFV6CaQqWZ04eBONjuE1jZhh_hXiLO8LBbrIl8Rfaee2xWaQVQB01gPpg0pgo3glQWabipTYhqGul1I68VzfDwtAqLhhbyp6EQ", "{\"kty\":\"RSA\",\"e\":\"AQAB\",\"use\":\"sig\",\"kid\":\"5cd61d45\",\"n\":\"nsAZboK0m7gaLW2qMUAxSTkXsaKCNAOe6csto5MJ9znklcOtmdL75YZHE60WeWhTvIdeiZlK-sDmW3PAwtOsTfohIyez2uKJ-DwRNhle7yFjUY9xiHXFYy7MRZA9zMmq8OOJIYsQsDurdClzRQGQj2T2N9mGcn09IaBE_DMMPdqXovHS97W5rHGFvTjX2I1ULOnn3HsIMJkayzRfD9kae98m-lDzpIhfH-ylYojRd0GXfgQMIhMmjdpjNyH9eSnqrn9LARbx6DfB2lvVVnpK-BB2TRtOuINmiEtADWXQvrfiYma6PcbDhi6SP0eah2hUgVQ5XLY2TzRuVT3gS9Fb5w\"}")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"name\": \"Max Mustermann\",\n    \"id\": \"1234567890\",\n    \"iat\": 1606428505\n}", payload)

	// RS384
	valid, payload, err = DecodeJwt("eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.nPsNEQxwlULiRrdNsMJ6trl7Nx9tCaZRZvCRHEK5JPxUkQDQPj0R0eubwUXhltir5-_Vbp-qbN5QZvvufeUpmJHZ5xVlCqQq268RWfQxLjXX_WBzlLPPxN-ciwOIcnM_x4gHdMW1Jg85LPSYJFgdiXDcq2-wVyxfmivQVQOi1Fp6BdZ5b5TXWQoxDws1h2GIVCV4rCff7gIxqHz-hyVSQvds12UDcvqAbD4iAHlTFocoUjqClfuANFNL2XhMP0UlJ6eRJUS28-ZQCw9CRv6BTnA1lnsEt2enO1oc6KGTf4J4n29uiHg-Ldib3cqHb4G6AMigOyngaJ-6NnqqonEUCw", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv\nvkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc\naT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy\ntvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0\ne+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb\nV6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9\nMwIDAQAB\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// RS512
	valid, payload, err = DecodeJwt("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.LFaSpaX3A_ShDt1ttid3Dnn2xq_8nQ5arVHQg8iCt5XjzSFCIW6DYfA8gYvHcZjHmqsv7KW78i029RQ8vhDYhEryIH7xL1Xn3PhWlg-oujJlhu8QVS91HGwEWZUrkFFrvLyJRk0vC16qMLgvgs7z-1lBXSHg2ssz6YQadtOeWOXrevet6rNeB8fOhs-pNqPxBhUg6577_81C9Nh5BJxARCIB__2TByF79jTPXQJ3wcO80dnQBnSI5CYUzTSV4br4VmASaZDYDmB8d3V7ASpNYZk6IJKVGvZC4keGAxmybBl24BRLw7f6CSzOw8a0fr6UX-LhOSzAfn_rynWnN7zn9A", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv\nvkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc\naT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy\ntvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0\ne+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb\nV6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9\nMwIDAQAB\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// ES256
	valid, payload, err = DecodeJwt("eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.emkpTcVeNAUhBIpKGJP7EXda73OZbzdnqCgr3Y2tq1eq-yc0-GSKp2_ISZXVitTfGrQ3_RTL_WR4kxQvwgYOlg", "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEVs/o5+uQbTjL3chynL4wXgUg2R9\nq9UU8I5mEovUf86QZ7kOBIjJwqnzD1omageEHWwHdBO6B+dFabmdT9POxg==\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// ES384
	valid, payload, err = DecodeJwt("eyJhbGciOiJFUzM4NCIsInR5cCI6IkpXVCIsImtpZCI6ImlUcVhYSTB6YkFuSkNLRGFvYmZoa00xZi02ck1TcFRmeVpNUnBfMnRLSTgifQ.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.MOhIpWkKvwpdzqYQVInGhByCn4bRPYflxtPLPPgeBQXaESF6gsfJLHN87ygDoea_XxlbEGBNgM6XFq-lYtcv8SFMuqyEnZDrOgtG3NiRMY6bnaLFhP2pBGQq2uIeyRkF", "-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEC1uWSXj2czCDwMTLWV5BFmwxdM6PX9p+\nPk9Yf9rIf374m5XP1U8q79dBhLSIuaojsvOT39UUcPJROSD1FqYLued0rXiooIii\n1D3jaW6pmGVJFhodzC31cy5sfOYotrzF\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// ES512
	valid, payload, err = DecodeJwt("eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCIsImtpZCI6InhaRGZacHJ5NFA5dlpQWnlHMmZOQlJqLTdMejVvbVZkbTd0SG9DZ1NOZlkifQ.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.AK2bdWotGzwg0UWr3c0wZJI3HjTha0MUN4YTKPGJYfpVFS_ZfPoC68xO95LIiAfIvLBCznbaY52Toq2o9Vmq5uwDATUJ_kpPiawG8nnKVc0hx5fiM22jQb-k_lmWolai767_eny-SCSgEhl-OohR2Odd_Me93ku4_96jda_0t92jgjL8", "-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBgc4HZz+/fBbC7lmEww0AO3NK9wVZ\nPDZ0VEnsaUFLEYpTzb90nITtJUcPUbvOsdZIZ1Q8fnbquAYgxXL5UgHMoywAib47\n6MkyyYgPk0BXZq3mq4zImTRNuaU9slj9TVJ3ScT3L1bXwVuPJDzpr5GOFpaj+WwM\nAl8G7CqwoJOsW7Kddns=\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// PS256
	valid, payload, err = DecodeJwt("eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.a1w4ktxG0ltBmWZJVNSMnRtPld7oRY3he_0zmSXXe5qI966VXsndCFRqaRv64H7qLSksKddD29C7acm5anVpF_qqU3elCUId_YSyKu_r6kU1IX2zzNeyUrEFcr00p07IT_0qXv54Wkz7mey34suGvktjgL5Yym6MO7yAPXwiVOfosKWwA0s1u1K3FQtZuPQ0hANn8QxqYDwnmJqF5ZHVx6PI4Vu_cMbAlncWd5iJ6dXzdLen7EhzKjZQxB720RRQ4GLm0sWTQk_oRSpLVOlTrvfGaRY3bSZHzcrBnbHJyokv2-xA7za4u1RkiDKg9GqKL-86msAol4O9ZQpAJfa1mw", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv\nvkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc\naT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy\ntvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0\ne+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb\nV6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9\nMwIDAQAB\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	// PS384
	valid, payload, err = DecodeJwt("eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.ATJTQALrd36Rn8mZiusWRrB8dFYbWnZiIccKdVy75RNIwsuj-rDnAjdprrpSO7nFkTHtGgffAfvmM-Zc31vC4IacN0wVBuAtzgtbdXJviQbKiuy2R0R4AiS_Ynpjdk5Z7NKB4AYbWlarfwbhGilbeg8xPcCmWFfb4JyA3zvYaGzPhn85YztgkOLndU3URWe6DcKUYOBHmy8gtMDOcEKHVU2HETG01TKMJeXmiRUwjhmfFn5VCzAuNEJvQHA_VIszTkZcyEhNQBLh610Yl1hVGbPvJXlLup0FxTj_dp8Sy2EYT2ASP-it2KiPXWkps0bEfA-81wkKIIFxustPcjPtqw", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv\nvkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc\naT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy\ntvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0\ne+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb\nV6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9\nMwIDAQAB\n-----END PUBLIC KEY-----")
	assert.Nil(t, err)
	assert.True(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	valid, payload, err = DecodeJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.TfVxPqlJ2xhpB5mKN6fVvb51Z_mIPbdNw2D_1b_tN0w", "wrongsecret")
	assert.Nil(t, err)
	assert.False(t, valid)
	assert.Equal(t, "{\n    \"id\": \"1234567890\",\n    \"name\": \"Max Mustermann\"\n}", payload)

	valid, _, err = DecodeJwt("invalid.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.TfVxPqlJ2xhpB5mKN6fVvb51Z_mIPbdNw2D_1b_tN0w", "mysecret")
	assert.NotNil(t, err)
	assert.False(t, valid)

	valid, _, err = DecodeJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.invalid.TfVxPqlJ2xhpB5mKN6fVvb51Z_mIPbdNw2D_1b_tN0w", "mysecret")
	assert.Nil(t, err)
	assert.False(t, valid)

	valid, _, err = DecodeJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiTWF4IE11c3Rlcm1hbm4ifQ.invalid", "mysecret")
	assert.Nil(t, err)
	assert.False(t, valid)

	valid, payload, err = DecodeJwt("eyJraWQiOiI1Y2Q2MWQ0NSIsInR5cCI6IkpXVCIsImFsZyI6IlJTMjU2In0.eyJuYW1lIjoiTWF4IE11c3Rlcm1hbm4iLCJpZCI6IjEyMzQ1Njc4OTAiLCJpYXQiOjE2MDY0Mjg1MDV9.U9WUtkDnNHe3SXGm1xJr8KfSZ9TvvnqmC3_r1akN5vHL2q73ZDf1ECjxNKj-Q6Oq4LJtJo9Hff0sbMsJXtLIfl0tvW4rB9gmdifu4KRj7QxfiCdUxdidaliYZXN_gCV10AIk8IvG7z_ITmbgd4-2hR4N-nqiSIsSfKmNcAyppvclXULFcw2_KqgefB1HQNasLiexH109YRWGTGIcSOS7OKEmjIeP-3Wd5kZJULJBWjryOJP6g-kNsFV6CaQqWZ04eBONjuE1jZhh_hXiLO8LBbrIl8Rfaee2xWaQVQB01gPpg0pgo3glQWabipTYhqGul1I68VzfDwtAqLhhbyp6EQ", "{\"kty\":\"RSA\",\"e\":\"AQAB\",\"use\":\"sig\",\"kid\":\"5cd61d45\",\"n\":\"invalid\"}")
	assert.Nil(t, err)
	assert.False(t, valid)
}
