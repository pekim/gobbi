char* open();
void close();
void call_function(int function_index);

typedef enum return_type {
    rt_int
} ReturnType;

typedef struct call_data {
    ReturnType return_type;
} CallData;
