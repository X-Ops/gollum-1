.. Autogenerated by Gollum RST generator (docs/generator/*.go)

RegExp
======

This filter rejects or accepts messages based on regular expressions.




Parameters
----------

**Expression**

  Messages matching this expression are passed on.
  This parameter is ignored when set to "". Expression is checked
  after ExpressionNot.
  By default this parameter is set to "".
  
  

**ExpressionNot**

  Messages *not* matching this expression are
  passed on. This parameter is ignored when set to "". ExpressionNot
  is checked before Expression.
  By default this parameter is set to "".
  
  

**ApplyTo**

  Defines which part of the message the filter is applied to.
  When set to "", this filter is applied to the message's payload. All
  other values denotes a metadata key.
  By default this parameter is set to "".
  
  

Parameters (from core.SimpleFilter)
-----------------------------------

**FilteredStream**

  This value defines the stream filtered messages get sent to.
  You can disable this behavior by setting the value to "".
  By default this parameter is set to "".
  
  

Examples
--------

This example accepts only accesslog entries with a return status of
2xx or 3xx not originated from staging systems.

.. code-block:: yaml

	 ExampleConsumer:
	   Type: consumer.Console
	   Streams: console
	   Modulators:
	     - filter.RegExp:
	       ExpressionNot: " stage\\."
	       Expression: "HTTP/1\\.1\\\" [23]\\d\\d"





