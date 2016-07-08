# Welcome to Sonic Pi v2.10

# chunky_png need to be installed under app/server/vendor
#  /Applications/Sonic Pi2.10.app/app/server/vendor/chunky_png-1.3.6
require 'chunky_png'

live_loop :readpng do
  begin
    image = ChunkyPNG::Image.from_file(File.expand_path("~/github/moere-telecoding/space/img/captured.png"))
    h = image.height
    w = image.width
    puts "image.height: #{h}"
    puts "image.width: #{w}"
    puts image.get_pixel(h/2, w/2) # get_pixel
    puts image[h/2, w/2] # get_pixel
    puts image[h/2, w/2].to_s(16).upcase # hex rgba

    sus = image[h/2, w/2]/0xFFFFFFFF.to_f*4*(rand+1)
    puts sus
    sample:ambi_piano,sustain:sus,rate:pitch_to_ratio([-5,-2,0,3].choose)*[1,2,4].choose*[1,-1].choose,start:rand
  rescue => e
    puts e.message
  end
  sleep 1
end
