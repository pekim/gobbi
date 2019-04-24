typedef enum return_type {
    rt_void,
    rt_int,
    rt_uint,
    rt_long,
    rt_double,
} ReturnType;

typedef struct call_data {
    ReturnType return_type;

    int value_int;
    unsigned int value_uint;
    long value_long;
    double value_double;
} CallData;

char* open();
void close();

void call_function(int function_index, CallData* data);
void *get_function(int function_index);
