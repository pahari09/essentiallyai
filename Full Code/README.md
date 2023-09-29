# Notes

 full code that decodes high frequency of data and with low latency in extremely robust and error tolerant manner while keeping the code maintainable.

 For high frequency and low latency:- configured kafka to utilize maximum cpu cores and memory with appropriate values.
 error tolerant- handled error wherever possible and continue to decode the data if few of them are unsucessful.
 maintainable code- separated the code wherever possible according to functionality. Wrote unit test as well.


 The code requires additional configurational changes to run successfully.
