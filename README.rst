========
tiktokgo
========

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/tiktokgo?status.svg
   :target: https://godoc.org/github.com/siongui/tiktokgo

.. image:: https://github.com/siongui/tiktokgo/workflows/Test%20Package/badge.svg
    :target: https://github.com/siongui/tiktokgo/blob/master/.github/workflows/build.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/tiktokgo
   :target: https://goreportcard.com/report/github.com/siongui/tiktokgo

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/tiktokgo/master/UNLICENSE


**tiktokgo** downloads user avatar photo and videos.

Currently this package can

- download user avatar photos and latest 5 video items (with watermark) without
  login if the user account is not private.
- download video item (with watermark) by the item URL.

Visit `example <tiktokdl/example/>` directory for examples.


Tested on:

  - `Ubuntu 20.10`_
  - `Go 1.16`_


Important Tricks
++++++++++++++++

According to README in `drawrowfly/tiktok-scraper`_, the video play/download
address is binded to the **tt_webid_v2** cookie value. To download the video
successfully, the same headers/cookies must be used both when access API or
metadata and when access/download videos.
The following three are necessary part (values copied from README in
`drawrowfly/tiktok-scraper`_ for illustration purpose):

.. code-block:: txt

  headers: {
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36",
    "referer": "https://www.tiktok.com/",
    "cookie": "tt_webid_v2=689854141086886123"
  },


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] | `github tiktok - Google search <https://www.google.com/search?q=github+tiktok>`_
       | `github tiktok - DuckDuckGo search <https://duckduckgo.com/?q=github+tiktok>`_
       | `github tiktok - Ecosia search <https://www.ecosia.org/search?q=github+tiktok>`_
       | `github tiktok - Qwant search <https://www.qwant.com/?q=github+tiktok>`_
       | `github tiktok - Bing search <https://www.bing.com/search?q=github+tiktok>`_
       | `github tiktok - Yahoo search <https://search.yahoo.com/search?p=github+tiktok>`_
       | `github tiktok - Baidu search <https://www.baidu.com/s?wd=github+tiktok>`_
       | `github tiktok - Yandex search <https://www.yandex.com/search/?text=github+tiktok>`_

.. _Go: https://golang.org/
.. _Ubuntu 20.10: https://releases.ubuntu.com/20.10/
.. _Go 1.16: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
.. _drawrowfly/tiktok-scraper: https://github.com/drawrowfly/tiktok-scraper
