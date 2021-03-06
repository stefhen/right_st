v1.4.4 / 2016-08-11
-------------------
* Fix bug, introduced in v1.4.0, where attachment names were getting mangled
* Correctly set default MCI for ServerTemplates on download
* Improve formatting of description fields when downloading ServerTemplates
* Default empty inputs to to "inherit from RightScript" rather than be a blank
  value when downloading ServerTemplates

v1.4.3 / 2016-08-11
-------------------
* Fix panic when using --mci-settings introduced in v1.4.2

v1.4.2 / 2016-08-11
-------------------
* Improve error messages, get rid of Description field showing up downloaded
  ServerTemplate yaml files for fixed revisions

v1.4.1 / 2016-08-08
-------------------
* Allow the update commands to work without a config file; this is useful if you
  need to use sudo to update since root probably does not have a config file.

v1.4.0 / 2016-08-08
-------------------
* Add `--script-path` argument to specify where to download RightScripts relative
  to the ServerTemplate download location ([#28])
* Fix a panic in ServerTemplate validate when the RightScript name is not known
  yet
* If multiple RightScripts in a ServerTemplate contain an attachment with the
  same name, download those attachments to different subdirectories ([#19])

[#28]: https://github.com/rightscale/right_st/pull/28
[#19]: https://github.com/rightscale/right_st/pull/19

v1.3.0 / 2016-07-15
-------------------
* Support for "published" or "external" MCIs. right_st will attempt to check for
  and use MultiCloudImages from the MultiCloud Marketplace if they exist ([#10])
* Support for full definition of MCIs in the ServerTemplate YAML file, including
  managing MCI settings and tags ([#11])
* Support specifying Packages field in RightScript comment metadata

[#10]: https://github.com/rightscale/right_st/issues/10
[#11]: https://github.com/rightscale/right_st/issues/11

v1.2.0 / 2016-06-05
-------------------
* Add ability to use "published" aka "external" RightScripts. These are
  RightScripts published in the MultiCloud Marketplace. These are linked to
  instead of downloaded to disk ([#16])

[#16]: https://github.com/rightscale/right_st/issues/16

v1.1.0 / 2016-04-15
-------------------
* Various fixes: Can specify a directory to download to. Better error messages.
* Now display script body in RightScript show command.
* Added --force option to scaffold to rescaffold an existing script [#5]
* Will clean up RightScript metadata on download if it is incorrect
* Attachments with the same name but belonging to different Rightscript will
  throw an error ([#9])
* Attachment names can now have absolute paths, or relative path components
  added in
* Fix [#13] with some trickery. Add a temporary MCI before deleting then adding
  the new MCI revision.
* Fix [#17] fully by only allowing actual host names.

[#5]: https://github.com/rightscale/right_st/issues/5
[#9]: https://github.com/rightscale/right_st/issues/9
[#13]: https://github.com/rightscale/right_st/issues/13
[#17]: https://github.com/rightscale/right_st/issues/17

v1.0.2 / 2016-04-07
-------------------
* Build with Go 1.6 on Travis CI.
* Fix [#14] by only doing an Inputs MultiUpdate when there are actually inputs
  to update.
* Actually unset inputs that are removed by getting the old inputs from a
  ServerTemplate before and setting any that are not being updated to "inherit".

[#14]: https://github.com/rightscale/right_st/issues/14

v1.0.1 / 2016-03-09
-------------------
* Fix a cosmetic bug where attachment downloads on Windows would display as
  `Downloading attachment into attachments/attachments\filename.txt` instead of
  `Downloading attachment into attachments/filename.txt`

v1.0.0 / 2016-03-03
-------------------
* Initial release
