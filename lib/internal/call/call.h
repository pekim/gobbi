typedef enum type {
    type_void,
    type_int,
    type_uint,
    type_long,
    type_double,
} Type;

typedef struct call_data {
    Type return_type;

    int value_int;
    unsigned int value_uint;
    long value_long;
    double value_double;
} CallData;

char* open();
void close();

void call_function(int function_index, CallData* data);
void *get_function(int function_index);
