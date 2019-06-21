|Build Status|

F5 BigIP Provider
=================

This provider allows management of F5 BigIP resources using Pulumi. This
provider uses the iControlREST API to perform management tasks, so it
will need to be installed and enabled on your F5 device before
proceeding.

Installing
----------

This package is available in many languages in the standard packaging
formats.

Node.js (Java/TypeScript)
~~~~~~~~~~~~~~~~~~~~~~~~~

To use from JavaScript or TypeScript in Node.js, install using either
``npm``:

::

   $ npm install @pulumi/f5bigip

or ``yarn``:

::

   $ yarn add @pulumi/f5bigip

Python
~~~~~~

To use from Python, install using ``pip``:

::

   $ pip install pulumi_f5bigip

Go
~~

To use from Go, use ``go get`` to grab the latest version of the library

::

   $ go get github.com/pulumi/pulumi-f5bigip/sdk/go/...

Reference
---------

For detailed reference documentation, please visit `the API
docs <https://pulumi.io/reference/pkg/nodejs/@pulumi/f5bigip/index.html>`__.

.. |Build Status| image:: https://travis-ci.com/pulumi/pulumi-f5bigip.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master
   :target: https://travis-ci.com/pulumi/pulumi-f5bigip
