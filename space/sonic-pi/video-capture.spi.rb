# Welcome to Sonic Pi v2.10

load "~/github/moere-telecoding/space/sonic-pi/libs/haskap-jam-pack/haskap-jam-p5/haskap-jam-p5.rb"
##| set_use_pty true
##| rp5_inline_sketch
rp5_inline_sketch({:full_screen => true})

load_library :video
include_package "processing.video"

def setup
  frame_rate 30
  smooth
  ##| size(640, 480)
  ##| size(1152, 648)
  size(1280, 720)
  ##| size displayWidth, displayHeight
  @video = Capture.new(self, width, height, 30)
  @video.start
end

def draw
  @video.read if @video.available?
  image(@video, 0, 0)

  if frame_count % (30 * 3) == 0
    puts(frame_rate)
    save(File.expand_path("~/github/moere-telecoding/space/img/captured.png"))
  end
end
