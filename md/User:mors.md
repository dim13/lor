# Сборка платформы Android \[draft\]

Это незаконченная статья. Сюда пока буду скидывать материалы

Platform:

`       Linux dmitry-laptop 2.6.35-23-generic #41-Ubuntu SMP Wed Nov 24 11:55:36 UTC 2010 x86_64 GNU/Linux`

`       repo init -u `<git://android.git.kernel.org/platform/manifest.git>` -b master #or android-2.2.1_r1`

  - for android-2.2\_r1.1 you need java1.5; for 2.3 -- java1.6
  - to switch from jdk 1.5 to 1.6 use sudo update-java-alternatives -s
    java-6-sun and restart terminal

PREPARE to build:

-----

`       ~/droid$ choosecombo`  
`       Build for the simulator or the device?`  
`            1. Device`  
`            2. Simulator`

`       Which would you like? [1] 1`

`       Build type choices are:`  
`            1. release`  
`            2. debug`

`       Which would you like? [1] 2`

`       Which product would you like? [sim] generic`

`       Variant choices are:`  
`            1. user`  
`            2. userdebug`  
`            3. eng`  
`       Which would you like? [eng] 3`

Command to make android platform:

`       m TARGET_ARCH=arm -j2 CROSS_COMPILE=/home/dmitry/droid/prebuilt/linux-x86/toolchain/arm-eabi-4.4.0/bin/arm-eabi-`

Use device, NOT simulator. Simulator isn't supported since cupcake. Or
you get an ERROR:

``       make: *** No rule to make target `out/debug/host/linux-x86/pr/sim/obj/lib/libc.so', ``  
``       needed by `out/debug/host/linux-x86/pr/sim/obj/SHARED_LIBRARIES/libexpat_intermediates/LINKED/libexpat.so'.  Stop.``

Kernel config:

`       ~/common$ ARCH=arm CROSS_COMPILE=../droid/prebuilt/linux-x86/toolchain/arm-eabi-4.2.1/bin/arm-eabi- make`

About building android' kernel:
<http://stackoverflow.com/questions/1809774/android-kernel-compile-and-test-with-android-emulator>

  - there is a prebuilt kernel

-----

\>\> 1) How to switch branches?

\> repo checkout otherbranch

\>\> 2) How to revert all changes?

\>repo sync -d \>repo abandon yourbranch

\>\> 3) How to get list of available branches?

\>repo branches

\>\> 4) How to get graph of branches?

\>repo forall -c 'gitk &' // Не делайте этого\!\!

(c)
<http://google-opensource.blogspot.com/2008/11/gerrit-and-repo-android-source.html>
Setup git proxy:
<http://discuz-android.blogspot.com/2008/10/repo-android-opensource-via-proxy.html>
