package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
	"fmt"
)

const Tag = 8

type Slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

//������Ƭ���ݽṹCreate
func (s *Slice) Create(l int, cp int, data ...int) {
	if s == nil {
		return
	}
	//�жϴ���
	if l < 0 || cp < 0 || l > cp || len(data) > cp {
		return
	}
	//�õ����ݵ�ַ
	s.Data = C.malloc(C.size_t(cp) * Tag)
	//��ָ���ת������ֵ���ͽ������
	p := uintptr(s.Data)
	for _, v := range data {
		*(*int)(unsafe.Pointer(p)) = v
		p += Tag
	}
	s.Len = l
	s.Cap = cp

}

//��ӡ��Ƭ
func (s *Slice) Print() {
	//�ݴ�
	if s == nil || s.Len == 0 {
		return
	}
	//�ѵ�ַת��������  Ȼ���ַ�Ϳ��Խ����������
	SliceData := uintptr(s.Data)
	//ѭ�������ݶ�����
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(* int)(unsafe.Pointer(SliceData)), " ")
		SliceData += Tag
	}
	fmt.Println()
}

//׷��Ԫ��Append(...����)
func (s *Slice) Append(Data ...int) {
	//�ݴ�
	if s == nil || len(Data) == 0 {
		return
	}
	//����������Ȼ���ܳ��ȳ������� ��Ҫ��������
	for s.Len+len(Data) > s.Cap {
		//��������˾ͽ���2������
		s.Cap = s.Cap * 2
	}

	//�ѵ�ַת��������,Ȼ���ַ�Ϳ��Խ���ƫ�� (���ݵļӺͼ�)
	SliceData := uintptr(s.Data)
	//�ѵ�ַƫ�Ƶ������
	SliceData += uintptr(s.Len) * Tag
	//ѭ����������� ��ȥ
	for _, v := range Data {
		*(* int)(unsafe.Pointer(SliceData)) = v
		SliceData += Tag
	}
	s.Len += len(Data)
}

//��ȡ��ƬԪ��GetData(�±�)int
func (s *Slice) GetData(Data int) int {
	//�ݴ�
	if s == nil || Data < 0 || Data >= s.Len {
		return -1
	}

	//����ַת��������
	SliceData := uintptr(s.Data)
	//�ҵ��±�ĵ�ַ
	spSliceData := SliceData + uintptr(Data)*Tag
	//�����ݻ�ת���ɵ�ַ ������ȥ
	return *(* int)(unsafe.Pointer(spSliceData))

}

//������ƬԪ��Search(����)�±�
func (s *Slice) Search(Data int) int {
	//�ݴ�
	if s == nil {
		return -1
	}
	//�ѵ�ַת��������
	SliceData := uintptr(s.Data)
	//ѭ�������ݱ���
	for i := 0; i < s.Len; i++ {
		if Data == *(*int)(unsafe.Pointer(SliceData)) {
			return i
		}
		SliceData += Tag
	}
	return -1
}

//����Ԫ�ص���Ƭ��insert(����,�±�)
func (s *Slice) insert(Data, index int) {
	//�ݴ�
	if s == nil || Data == 0 || index > s.Len || index < 0 {
		return
	}
	if s.Len+1 > s.Cap {
		s.Cap = s.Cap * 2
	}
	//�������ݱ���
	// ��ַ -- ��ֵ
	p := uintptr(s.Data)

	// ���� index ��Ӧ���ڴ��ַ
	p += uintptr(index) * Tag

	// ������� tmp , �����ƶ� index ����Ԫ��.
	tmp := uintptr(s.Data) + uintptr(s.Len)*Tag // �� tmp ָ����ƬβԪ�صĺ�һ��λ��.

	// �Ӻ���ǰѭ��,ǰһ��Ԫ��,����һ��Ԫ�ظ�ֵ.
	for i := s.Len; i > index; i-- {
		// ǰһ��Ԫ��,����һ��Ԫ�ظ�ֵ
		*(*int)(unsafe.Pointer(tmp)) = *(*int)(unsafe.Pointer(tmp - 8))
		tmp -= Tag // ָ��ǰ��
	}

	// ��index λ��д��,�²�������ֵ
	*(*int)(unsafe.Pointer(p)) = Data

	// ������Ч����
	s.Len++
}

//ɾ����Ƭ�е�Ԫ��Delete(�±�)
func (s *Slice) Delete(index int) {
	//�ݴ�
	if s == nil || index > s.Len || index < 0 {
		return
	}
	//���е�ַ������ת��
	SliceData := uintptr(s.Data)
	//SliceData += Tag * uintptr(index)
	//����
	ptrindex := s.Len - index - 1
	for i := 0; i < ptrindex; i++ {
		if i==ptrindex {
			*(*int)(unsafe.Pointer(SliceData + Tag))=0
		}
		*(*int)(unsafe.Pointer(SliceData)) = *(*int)(unsafe.Pointer(SliceData + Tag))
		SliceData+=Tag
	}
	s.Len--
}

//������Ƭ
func (s *Slice)Destroy()  {
	s.Len=0
	s.Data=nil
	s.Cap=0
}

func main() {
	var s Slice
	s.Create(5, 5, 1, 2, 3, 4, 5)
	//s.Append(123, 123, 12, 2, 2, 2, 2, 2)
	//s.Print()
	//fmt.Println(s.GetData(11))

	s.Delete(0)
	s.Print()
	s.Destroy()

	s.Print()
}
