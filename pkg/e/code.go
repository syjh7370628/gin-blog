package e

const (
	SUCCESS               = 200
	ERROR                 = 500
	INVALID_PARAMS        = 400
	INVALID_JSON_PARAMS   = 422
	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_NOT_EXIST_COMMENT        = 10031
	ERROR_CHECK_EXIST_COMMENT_FAIL = 10032
	ERROR_ADD_COMMENT_FAIL         = 10033
	ERROR_DELETE_COMMENT_FAIL      = 10034
	ERROR_EDIT_COMMENT_FAIL        = 10035
	ERROR_COUNT_COMMENT_FAIL       = 10036
	ERROR_GET_COMMENTS_FAIL        = 10037
	ERROR_GET_COMMENT_FAIL         = 10038
	ERROR_GEN_COMMENT_POSTER_FAIL  = 10039

	ERROR_NOT_EXIST_CATEGORY        = 10051
	ERROR_CHECK_EXIST_CATEGORY_FAIL = 10052
	ERROR_ADD_CATEGORY_FAIL         = 10053
	ERROR_DELETE_CATEGORY_FAIL      = 10054
	ERROR_EDIT_CATEGORY_FAIL        = 10055
	ERROR_COUNT_CATEGORY_FAIL       = 10056
	ERROR_GET_CATEGORYS_FAIL        = 10057
	ERROR_GET_CATEGORY_FAIL         = 10058
	ERROR_GEN_CATEGORY_POSTER_FAIL  = 10059
	ERROR_ALREADY_EXIST_CATEGORY    = 10060

	ERROR_NOT_EXIST_MESSAGE        = 10071
	ERROR_CHECK_EXIST_MESSAGE_FAIL = 10072
	ERROR_ADD_MESSAGE_FAIL         = 10073
	ERROR_DELETE_MESSAGE_FAIL      = 10074
	ERROR_EDIT_MESSAGE_FAIL        = 10075
	ERROR_COUNT_MESSAGE_FAIL       = 10076
	ERROR_GET_MESSAGES_FAIL        = 10077
	ERROR_GET_MESSAGE_FAIL         = 10078
	ERROR_GEN_MESSAGE_POSTER_FAIL  = 10079

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_CAPTCHA_QUERY_PARAM_ERROR  = 20011
	ERROR_CAPTCHA_RESOURCE_NOT_FOUNT = 20012
	ERROR_CAPTCHA_SYSTEM_ERROE       = 20013
	ERROR_CAPTCHA_ERROR              = 20014

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)
