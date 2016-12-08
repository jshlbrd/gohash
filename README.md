# gohash
gohash is a simple directory scanning and MD5 hashing package written in Go. It has options to specify the input directory (files to be scanned), output file (where the output shown below goes), and maximum file size for hashed files (so you don't overstress your computer hashing enormous files).

## Command execution

<pre><code>XXXXXXXXXXX$ ./bin/gohash -h
Usage of ./bin/gohash:
  -maxSize int
    	The maximum size of files to be hashed (default 10000000)
  -outputFile string
    	The file where MD5 hashes will be output (default "/tmp/gohash.dat")
  -scanDir string
    	The directory that will be scanned</pre></code>

## Output

<pre><code>/usr/local/bro/bin/bro: e3e65b3b6cae4f2255db11646b0aebc3
/usr/local/bro/bin/bro-config: ae41f9132ab2a442a159ea5cce426dd1
/usr/local/bro/bin/bro-cut: fe1d3232ffbf413dd2a4928bf6ab98cb
/usr/local/bro/bin/broccoli-config: 88ecca823ca13c1a595d732cd21d6396
/usr/local/bro/bin/broctl: 406b2cf55b3a0150bead3fa6c61dfebf
/usr/local/bro/bin/capstats: d503f8952a5b12e3322d13733c1f2eb9
/usr/local/bro/bin/trace-summary: 531cc127cf73a5804a4b4d17f393c4a9
/usr/local/bro/etc/broccoli.conf: b7fac220821be497b0d5a79f695c01a4
/usr/local/bro/etc/broctl.cfg: eecf5b01dea7d569370af3271e1c0f15
/usr/local/bro/etc/broctl.cfg.example: cdd476c88cc9f2a0a2e6f2966132e59e
/usr/local/bro/etc/networks.cfg: afed5d6ba219183dd5d7107b60d9ad40
/usr/local/bro/etc/node.cfg: 91e7e9acab1862487f22d9f3c15ece02
/usr/local/bro/etc/node.cfg.example: 754db8253db4ebd1262762895fefd5ca
/usr/local/bro/include/broccoli.h: e56d5b44f979d2693a0ae03e721f60b6
/usr/local/bro/lib/broctl/BroControl/__init__.py: 7759f50933068da38483420a7e9988a4
/usr/local/bro/lib/broctl/BroControl/brocmd.py: dfffbc6bbe26729ee332d2cd72ac2706
/usr/local/bro/lib/broctl/BroControl/broctl.py: ff8cb14bdd566ca8b70e6ceafad9552f
/usr/local/bro/lib/broctl/BroControl/cmdresult.py: 7f5938c7a25ef75759cdbdf8c92ec0bf
/usr/local/bro/lib/broctl/BroControl/config.py: 2e67a1fe00ed35e2dd7f1d6a23d4a0e8
/usr/local/bro/lib/broctl/BroControl/control.py: 5e3fac66776e21bb9ebc2ca777058d1e
/usr/local/bro/lib/broctl/BroControl/cron.py: 680305ed2422e00ca1e2c21f691e3b5f
/usr/local/bro/lib/broctl/BroControl/doc.py: 030fa15f89ca2483949f8de83fa935b5
/usr/local/bro/lib/broctl/BroControl/events.py: 4d5053f9f0d1a9c9490a1043a7363bf6
/usr/local/bro/lib/broctl/BroControl/execute.py: 82dfef0862f371cf039b6e93ce414ded
/usr/local/bro/lib/broctl/BroControl/install.py: b88e2858a4555a3a9fa3dbd2f40bb637
/usr/local/bro/lib/broctl/BroControl/node.py: 080692c8172ad1a7a9839930f738828c
/usr/local/bro/lib/broctl/BroControl/options.py: 104d6ec57ca3db6021e69aa70c822cc1
/usr/local/bro/lib/broctl/BroControl/plugin.py: d522ab8335c67b90799623483f151d10
/usr/local/bro/lib/broctl/BroControl/pluginreg.py: bb70e266f624b4c484bd4e30497b3fb9
/usr/local/bro/lib/broctl/BroControl/printdoc.py: ae8104a096ffb3b6627d3da5a990cc92
/usr/local/bro/lib/broctl/BroControl/py3bro.py: e9c4119e75f95fa64306e865fff5b008
/usr/local/bro/lib/broctl/BroControl/ssh_runner.py: 8717ff5f209ad6f196caeba5748242fd
...</pre></code>
