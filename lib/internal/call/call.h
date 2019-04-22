typedef enum return_type {
    rt_int
} ReturnType;

typedef struct call_data {
    ReturnType return_type;
} CallData;

char* open();
void close();

void call_function(int function_index, CallData* data);
void *get_function(int function_index);
