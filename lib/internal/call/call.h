typedef enum return_type {
    rt_void,
    rt_int,
    rt_uint,
} ReturnType;

typedef struct call_data {
    ReturnType return_type;
    int return_int;
    uint return_uint;
} CallData;

char* open();
void close();

void call_function(int function_index, CallData* data);
void *get_function(int function_index);
