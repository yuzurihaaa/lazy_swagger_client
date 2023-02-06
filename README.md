# Lazy Swagger Client
1. A lazy, just want to get things done fast without having to write again logics for api request that is already defined in swagger files.
2. No type checking, no auto generate code - just getting things done without code generation hassle.

# Motivation
1. I'm trying to build something like https://github.com/swagger-api/swagger-js for go. The whole idea is to just call the endpoint without having problem dealing with types (request / response). 
If you want something that is type-safety, use other libraries.
2. I don't want to deal with code generation. Let me shoot my foot by my own mistake, not due to some code generation that I'll need to modify swagger files to satisfy code generator.

# How-to use
1. Initialize swagger client.
    ```go
    swagger := lazy_swagger.NewSwaggerF("path-to-swagger.json")
    ```
2. Execute the function that you want to call.
   ```go
    res, err := c.swagger.Execute(toCtx, "Unique_swagger_operation_id", lazy_swagger.Args{})
    if err != nil {
        // handle request error here
	}
   if res.StatusCode != http.StatusOK {
        // handle response here.
   }
   ```
    
