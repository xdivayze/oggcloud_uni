import CommonFrame from "./CommonFrame";

export default function ExpandedNavbar({
  height,
  showText: textSize,
  additional
}: {
  height: string;
  showText: string;
  additional: string
}) {
  var classNames = `w-full justify-center flex flex-row  relative ${height} ${textSize} ${additional}`;
  return (
    <nav className={classNames}>
      <CommonFrame
        text="OGGLabs Interactive Research Experience"
        to="/about"
        textSize={textSize}
        color="bg-blue-ogg"
        margin="ml-18"
        z="z-5"
      />
      <CommonFrame
        text="Home"
        textSize={textSize}
        to="/"
        color="bg-teal-ogg-1"
        margin="-ml-7 -pl-7 "
        z="z-4"
      />
      <CommonFrame
        text="Register"
        to="/register"
        textSize={textSize}
        color="bg-green-ogg-2"
        margin="-ml-7 -pl-7"
        z="z-3"
      />
      <CommonFrame
        text="Login"
        to="/login"
        textSize={textSize}
        color="bg-brown-ogg"
        margin="-ml-7 -pl-7"
        z="z-2"
      />
      <CommonFrame
        text="Help"
        to="/help"
        textSize={textSize}
        color="bg-teal-ogg"
        margin="-ml-7 -pl-7"
        z="z-1"
      />
    </nav>
  );
}
