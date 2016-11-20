package classpath
import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry  {
	baseDir := path[len(path) - 1]
	compositeEntry := []Entry{}
	walfFn := func(path string,info os.FileInfo,err error) {
		if err != nil{
			return err
		}
		if info.IsDir() && path != baseDir{
			return filepath.SkipDir
		}
		if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR"){
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir,walfFn)
	return compositeEntry
}