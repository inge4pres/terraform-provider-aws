	"strconv"
// ParseMultiFileDiff parses a multi-file unified diff. It returns an error if
// parsing failed as a whole, but does its best to parse as many files in the
// case of per-file errors. If it cannot detect when the diff of the next file
// begins, the hunks are added to the FileDiff of the previous file.
	if err != nil && err != io.EOF {
	unquotedOrigName, err := strconv.Unquote(origName)
	if err == nil {
		origName = unquotedOrigName
	}
	unquotedNewName, err := strconv.Unquote(newName)
	if err == nil {
		newName = unquotedNewName
	}

	var err error
	lineCount := len(fd.Extended)
	if lineCount > 0 && !strings.HasPrefix(fd.Extended[0], "diff --git ") {
		return false
	}
	case (lineCount == 3 || lineCount == 4 && strings.HasPrefix(fd.Extended[3], "Binary files ") || lineCount > 4 && strings.HasPrefix(fd.Extended[3], "GIT binary patch")) &&
		strings.HasPrefix(fd.Extended[1], "new file mode "):
		fd.NewName, err = strconv.Unquote(names[1])
		if err != nil {
			fd.NewName = names[1]
		}
	case (lineCount == 3 || lineCount == 4 && strings.HasPrefix(fd.Extended[3], "Binary files ") || lineCount > 4 && strings.HasPrefix(fd.Extended[3], "GIT binary patch")) &&
		strings.HasPrefix(fd.Extended[1], "deleted file mode "):
		fd.OrigName, err = strconv.Unquote(names[0])
		if err != nil {
			fd.OrigName = names[0]
		}
	case lineCount == 4 && strings.HasPrefix(fd.Extended[2], "rename from ") && strings.HasPrefix(fd.Extended[3], "rename to "):
		fd.OrigName, err = strconv.Unquote(names[0])
		if err != nil {
			fd.OrigName = names[0]
		}
		fd.NewName, err = strconv.Unquote(names[1])
		if err != nil {
			fd.NewName = names[1]
		}
	case lineCount == 3 && strings.HasPrefix(fd.Extended[2], "Binary files ") || lineCount > 3 && strings.HasPrefix(fd.Extended[2], "GIT binary patch"):
		fd.OrigName, err = strconv.Unquote(names[0])
		if err != nil {
			fd.OrigName = names[0]
		}
		fd.NewName, err = strconv.Unquote(names[1])
		if err != nil {
			fd.NewName = names[1]
		}