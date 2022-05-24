Go Logitech Litra Glow Driver
=============================

This GO module implements a basic way to interact with a Logitech Litra Glow
light. It uses `karalabe/usb <https://github.com/karalabe/usb>`_ for the low
level USB communications. I have only tested this on Linux.

The reverse engineered USB protocol comes from the `kharyam/litra-driver
<https://github.com/kharyam/litra-driver>`_ Python implementation.

UDEV Config
-----------

Before you can use the light on Linux as a non-root user, you need to set up
the following UDEV rule in ``/etc/udev/rules.d/82-litra-glow.rules``::

	SUBSYSTEM=="usb", ATTRS{idVendor}=="046d", ATTRS{idProduct}=="c900", MODE:="0666", GROUP="plugdev"

Then restart UDEV to refresh it rules::

	udevadm control --reload-rules && udevadm trigger

And then (re-)plugin your light again.

Usage
-----

Obtain the driver with::

	go get derickr/go-litra-driver

Import the driver in your ``.go`` file::

	import derickr/go-litra-driver

In order to use the driver to control the light, create a new instance of the
``LitraDevice`` struct::

	ld, err := litra.New()

Currently, the driver only supports one light. If the driver can't open the
USD device, an error will be returned.

You can then use the ``ld`` variable as a handle to issue control statements.

================  =============================  =====================================
Task              Method                         Arguments
----------------  -----------------------------  -------------------------------------
Turn Light On     ld.TurnOn()
Turn Light Off    ld.TurnOff()
Set Brightness    ld.SetBrightness(level int)    ``level`` is brightness from 0 to 100
Set Temperature   ld.SetTemperature(temp int16)  ``temp`` is light temperature
                                                 from 2700K to 6500K
Close Connection  ld.Close()
================  =============================  =====================================

If you supply an out-of-range value to ``SetBrightness`` or ``SetTemperature``
it clamps to the supported range.
