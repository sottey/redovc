desc "Builds redovc for release"

Envs = [
  { goos: "darwin", arch: "386" },
  { goos: "darwin", arch: "amd64" },
  { goos: "linux", arch: "arm" },
  { goos: "linux", arch: "arm64" },
  { goos: "linux", arch: "386" },
  { goos: "linux", arch: "amd64" },
  { goos: "windows", arch: "386" },
  { goos: "windows", arch: "amd64" }
].freeze

Version = "1.7.3".freeze

task :build do
  `rm -rf dist/#{Version}`
  Envs.each do |env|
    ENV["GOOS"] = env[:goos]
    ENV["GOARCH"] = env[:arch]
    puts "Building #{env[:goos]} #{env[:arch]}"
    `GOOS=#{env[:goos]} GOARCH=#{env[:arch]} CGO_ENABLED=0 go build -v -o dist/#{Version}/redovc`
    if env[:goos] == "windows"
      puts "Creating windows executable"
      `mv dist/#{Version}/redovc dist/#{Version}/redovc.exe`
      `cd dist/#{Version} && zip redovc_win.zip redovc.exe`
      puts "Removing windows executable"
      `rm -rf dist/#{Version}/redovc.exe`
    else
      puts "Tarring #{env[:goos]} #{env[:arch]}"
      `cd dist/#{Version} && tar -czvf redovc#{env[:goos]}_#{env[:arch]}.tar.gz redovc`
      puts "Removing dist/#{Version}/redovc"
      `rm -rf dist/#{Version}/redovc`
    end
  end
end

desc "Tests all the things"
task :test do
  system "go test ./..."
end

task default: :test
