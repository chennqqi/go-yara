#include <yara.h>
#include <stdio.h>


int main()
{
	yr_initialize();

	printf("yara initialize ok, version=%s\n", YR_VERSION);

	yr_finalize();

	return 0;	
}

