#### ImageMagick 

```shell
identify可以展示图像的详细信息
identify -verbose   1627138076033.jpeg

//jpeg 转 pdf 
convert -density 72 qq.jpeg qqq.pdf 
//pdf 转png 72为 分辨率的百分比
convert -density 72  1627138076033.jpeg  1627138076033_2.jpeg
//缩略图生成 30为 质量的百分比，宽高不变
convert -quality 30  1627138076033.jpeg  1627138076033_2.jpeg
//缩略图生成 30为 质量的百分比，按指定的宽度缩放，保持宽高比例
convert -quality 30 -resize 500 1627138076033.jpeg  1627138076033_2.jpeg
 //缩略图生成 30为质量的百分比，按指定高度缩放，保持宽高比
convert -quality 30 -resize 300 1627138076033.jpeg  1627138076033_2.jpeg
convert -quality 100 -resize 300 1627138076033.jpeg  1627138076033_2.jpeg
  
  
convert -density 30  test.gif  test_1.gif
convert -quality 30  test.gif  test_1.gif

背景可能会有轻微的抖动，同时，视频的每一帧画面采用了有损压缩，相近的颜色会被压缩为同一种颜色，因此同样的颜色，在不同的帧里面可能会被近似成不同的颜色。这两个因素就导致透明度优化效果很差。由于透明度优化是一个非常复杂的过程，牵涉到LZW压缩算法，因此这种情况下，优化可能反而使得LZW压缩的选择变差。所以最终优化过的gif，大小就和原图差不多，甚至更大
优化的时候，将相近的颜色视为相同的颜色，这样就可以抵消抖动和色彩压缩的影响。这通过设置fuzz因子来达成：
-fuzz选择多少能取得最大的压缩效果，同时对画质的影响可以接受
convert test.gif -fuzz 15% -layers Optimize result.gif
convert test.gif -fuzz 5% -layers Optimize result.gif

压缩bmp文件
convert -quality 100 -resize 300 testbmp.bmp  testbmp_1.jpeg
flop:水平翻转
convert -flop  -quality 100 -resize 300 1627138076033.jpeg  1627138076033_2.jpeg
flip:垂直翻转
convert -flip   -quality 100 -resize 300 1627138076033.jpeg  1627138076033_2.jpeg
-rotate:指定旋转的度数
顺时针旋转90度
convert  -rotate 90  -quality 100 -resize 300 1627138076033.jpeg  1627138076033_2.jpeg
逆时针旋转90度
convert  -rotate -90  -quality 100 -resize 300 1627138076033.jpeg  1627138076033_2.jpeg

```

