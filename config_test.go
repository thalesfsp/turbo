package turbo

import (
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getFileRootPath() string {
	p, e := filepath.Abs("../../..")
	if e != nil {
		panic(e)
	}
	return p
}

func TestConfig(t *testing.T) {
	c := NewConfig("grpc", "test/service_test.yaml")
	c.configs[fileRootPath] = getFileRootPath()
	assert.Equal(t, "production", c.Env())
	assert.Equal(t, "grpc", RpcType)
	assert.Equal(t, getFileRootPath()+"/github.com/vaporz/turbo/test", c.ServiceRootPath())

	assert.Equal(t, int64(8081), c.HTTPPort())

	assert.Equal(t, "YourService", c.GrpcServiceNames()[0])
	c.configs[grpcServiceName] = "test"
	assert.Equal(t, "test", c.GrpcServiceNames()[0])

	assert.Equal(t, "127.0.0.1", c.GrpcServiceHost())
	assert.Equal(t, "50051", c.GrpcServicePort())

	assert.Equal(t, "YourService", c.ThriftServiceNames()[0])
	c.configs[thriftServiceName] = "test thrift"
	assert.Equal(t, "test thrift", c.ThriftServiceNames()[0])

	assert.Equal(t, "127.0.0.1", c.ThriftServiceHost())
	assert.Equal(t, "50052", c.ThriftServicePort())

	assert.Equal(t, true, c.FilterProtoJson())
	c.configs[filterProtoJson] = strconv.FormatBool(false)
	assert.Equal(t, false, c.FilterProtoJson())

	assert.Equal(t, false, c.FilterProtoJsonInt64AsNumber())
	c.configs[filterProtoJsonInt64AsNumber] = strconv.FormatBool(true)
	assert.Equal(t, false, c.FilterProtoJsonInt64AsNumber())
	c.configs[filterProtoJson] = strconv.FormatBool(true)
	assert.Equal(t, true, c.FilterProtoJsonInt64AsNumber())
	c.configs[filterProtoJsonInt64AsNumber] = strconv.FormatBool(false)
	assert.Equal(t, false, c.FilterProtoJsonInt64AsNumber())

	c.configs[filterProtoJson] = strconv.FormatBool(false)
	assert.Equal(t, false, c.FilterProtoJsonEmitZeroValues())
	c.configs[filterProtoJsonEmitZeroValues] = strconv.FormatBool(true)
	assert.Equal(t, false, c.FilterProtoJsonEmitZeroValues())
	c.configs[filterProtoJson] = strconv.FormatBool(true)
	assert.Equal(t, true, c.FilterProtoJsonEmitZeroValues())
	c.configs[filterProtoJsonEmitZeroValues] = strconv.FormatBool(false)
	assert.Equal(t, false, c.FilterProtoJsonEmitZeroValues())

	assert.Equal(t, "GET,POST", c.mappings[urlServiceMaps][0][0])
	assert.Equal(t, "/hello", c.mappings[urlServiceMaps][0][1])
	assert.Equal(t, "SayHello", c.mappings[urlServiceMaps][0][2])
	assert.Equal(t, "GET", c.mappings[urlServiceMaps][1][0])
	assert.Equal(t, "/eat_apple/{num:[0-9]+}", c.mappings[urlServiceMaps][1][1])
	assert.Equal(t, "EatApple", c.mappings[urlServiceMaps][1][2])
	assert.Equal(t, "LogInterceptor", c.mappings[interceptors][0][2])
	assert.Equal(t, "preprocessor", c.mappings[preprocessors][0][2])
	assert.Equal(t, "postprocessor", c.mappings[postprocessors][0][2])
	assert.Equal(t, "hijacker", c.mappings[hijackers][0][2])
	assert.Equal(t, "convertor", c.mappings[convertors][0][1])
	assert.Equal(t, "error_handler", c.ErrorHandler())

	c.loadFieldMapping()
	assert.Equal(t, "CommonValues values", c.fieldMappings["SayHelloRequest"][0])
}

func TestHttpPortPanic(t *testing.T) {
	c := NewConfig("grpc", "test/service_test.yaml")
	p := c.HTTPPort()
	defer func() {
		c.configs[httpPort] = strconv.FormatInt(p, 10)
		if err := recover(); err != nil {
			assert.Equal(t, "[http_port] is required!", err)
		} else {
			t.Errorf("The code did not panic")
		}
	}()
	c.configs[httpPort] = ""
	c.HTTPPort()
}
