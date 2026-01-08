#include <libevdev/libevdev-uinput.h>
#include <libevdev/libevdev.h>
#include <linux/input-event-codes.h>
#include <linux/input.h>
#include <sys/types.h>
#include <unistd.h>
struct libevdev_uinput *create_controller() {

  struct libevdev *dev = libevdev_new();
  libevdev_set_name(dev, "DSU");
  libevdev_set_id_bustype(dev, BUS_USB);

  libevdev_set_id_vendor(dev, 0x045e);
  //  libevdev_set_id_vendor(dev, 0x28de);
  libevdev_set_id_product(dev, 0x028e);
  //  libevdev_set_id_product(dev, 0xffff);
  libevdev_set_id_version(dev, 0x0100);
  //  libevdev_set_id_version(dev, 0x001);
  libevdev_enable_event_type(dev, EV_KEY);

  libevdev_enable_event_code(dev, EV_KEY, BTN_SOUTH, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_EAST, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_WEST, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_NORTH, NULL);
    libevdev_enable_event_code(dev, EV_KEY, BTN_TL, NULL);
    libevdev_enable_event_code(dev, EV_KEY, BTN_TR, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_SELECT, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_START, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_THUMBL, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_THUMBR, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_DPAD_DOWN, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_DPAD_LEFT, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_DPAD_UP, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_DPAD_RIGHT, NULL);
  libevdev_enable_event_code(dev, EV_KEY, BTN_MODE, NULL);

  struct input_absinfo stick = {
      .minimum = -32768, .maximum = 32767, .flat = 128, .fuzz = 0};
       struct input_absinfo trigger = {
           .minimum = 0, .maximum = 255, .flat = 0, .fuzz = 0};

  libevdev_enable_event_type(dev, EV_ABS);
  libevdev_enable_event_code(dev, EV_ABS, ABS_X, &stick);
  libevdev_enable_event_code(dev, EV_ABS, ABS_Y, &stick);
  libevdev_enable_event_code(dev, EV_ABS, ABS_RX, &stick);
  libevdev_enable_event_code(dev, EV_ABS, ABS_RY, &stick);


  libevdev_enable_event_code(dev, EV_ABS, ABS_Z, &trigger);
  libevdev_enable_event_code(dev, EV_ABS, ABS_RZ, &trigger);

  struct libevdev_uinput *uidev;
  libevdev_uinput_create_from_device(dev, LIBEVDEV_UINPUT_OPEN_MANAGED, &uidev);

  return uidev;
}
void destroy_device(struct libevdev_uinput *uidev) {
  libevdev_uinput_destroy(uidev);
}
void btn_down(struct libevdev_uinput *uidev, unsigned int code) {
  libevdev_uinput_write_event(uidev, EV_KEY, code, 1);
}
void btn_up(struct libevdev_uinput *uidev, unsigned int code) {
  libevdev_uinput_write_event(uidev, EV_KEY, code, 0);
}
void syn(struct libevdev_uinput *uidev) {
  libevdev_uinput_write_event(uidev, EV_SYN, SYN_REPORT, 0);
}

void set_axis(struct libevdev_uinput *uidev, unsigned int code,
              unsigned int value) {
  libevdev_uinput_write_event(uidev, EV_ABS, code, value);
}
